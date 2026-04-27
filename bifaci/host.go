package bifaci

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"sort"
	"sync"

	"github.com/machinefabric/capdag-go/urn"
)

// CapHandler is a function that handles a peer invoke request.
// It receives the concatenated payload bytes and returns response bytes.
type CapHandler func(payload []byte) ([]byte, error)

// ResponseChunk represents a response chunk from a cartridge (matches Rust ResponseChunk)
type ResponseChunk struct {
	Payload []byte
	Seq     uint64
	Offset  *uint64
	Len     *uint64
	IsEof   bool
}

// CartridgeResponseType indicates whether a response is single or streaming
type CartridgeResponseType int

const (
	CartridgeResponseTypeSingle CartridgeResponseType = iota
	CartridgeResponseTypeStreaming
)

// CartridgeResponse represents a complete response from a cartridge
type CartridgeResponse struct {
	Type      CartridgeResponseType
	Single    []byte
	Streaming []*ResponseChunk
}

// FinalPayload gets the final payload
func (pr *CartridgeResponse) FinalPayload() []byte {
	switch pr.Type {
	case CartridgeResponseTypeSingle:
		return pr.Single
	case CartridgeResponseTypeStreaming:
		if len(pr.Streaming) > 0 {
			return pr.Streaming[len(pr.Streaming)-1].Payload
		}
		return nil
	default:
		return nil
	}
}

// Concatenated concatenates all payloads into a single buffer
func (pr *CartridgeResponse) Concatenated() []byte {
	switch pr.Type {
	case CartridgeResponseTypeSingle:
		result := make([]byte, len(pr.Single))
		copy(result, pr.Single)
		return result
	case CartridgeResponseTypeStreaming:
		totalLen := 0
		for _, chunk := range pr.Streaming {
			totalLen += len(chunk.Payload)
		}
		result := make([]byte, 0, totalLen)
		for _, chunk := range pr.Streaming {
			result = append(result, chunk.Payload...)
		}
		return result
	default:
		return nil
	}
}

// HostError represents errors from the cartridge host
type HostError struct {
	Type    HostErrorType
	Message string
	Code    string
}

type HostErrorType int

const (
	HostErrorTypeCbor HostErrorType = iota
	HostErrorTypeIo
	HostErrorTypeCartridgeError
	HostErrorTypeUnexpectedFrameType
	HostErrorTypeProcessExited
	HostErrorTypeHandshake
	HostErrorTypeClosed
	HostErrorTypeSendError
	HostErrorTypeRecvError
	HostErrorTypePeerInvokeNotSupported
)

func (e *HostError) Error() string {
	switch e.Type {
	case HostErrorTypeCbor:
		return fmt.Sprintf("CBOR error: %s", e.Message)
	case HostErrorTypeIo:
		return fmt.Sprintf("I/O error: %s", e.Message)
	case HostErrorTypeCartridgeError:
		return fmt.Sprintf("Cartridge returned error: [%s] %s", e.Code, e.Message)
	case HostErrorTypeUnexpectedFrameType:
		return fmt.Sprintf("Unexpected frame type: %s", e.Message)
	case HostErrorTypeProcessExited:
		return "Cartridge process exited unexpectedly"
	case HostErrorTypeHandshake:
		return fmt.Sprintf("Handshake failed: %s", e.Message)
	case HostErrorTypeClosed:
		return "Host is closed"
	case HostErrorTypeSendError:
		return "Send error: channel closed"
	case HostErrorTypeRecvError:
		return "Receive error: channel closed"
	case HostErrorTypePeerInvokeNotSupported:
		return fmt.Sprintf("Peer invoke not supported: %s", e.Message)
	default:
		return fmt.Sprintf("Unknown error: %s", e.Message)
	}
}

// =========================================================================
// Multi-cartridge host
// =========================================================================

// cartridgeEvent is an internal event from a cartridge reader goroutine.
type cartridgeEvent struct {
	cartridgeIdx int
	frame        *Frame
	isDeath      bool
}

// capTableEntry maps a cap URN to a cartridge index.
type capTableEntry struct {
	capUrn       string
	cartridgeIdx int
}

// rxidKey composes the (XID, RID) tuple used to route incoming
// requests from the relay. Mirrors the Rust host's
// `incoming_rxids: HashMap<(MessageId, MessageId), usize>` — XID
// (routing_id, assigned by the RelaySwitch) plus RID (the
// engine-side request id) together identify a request body
// uniquely. Composing them as a single string lets us use Go's
// map[string] without a separate hash impl.
type rxidKey struct {
	xid string
	rid string
}

func makeRxidKey(xid, rid MessageId) rxidKey {
	return rxidKey{xid: xid.ToString(), rid: rid.ToString()}
}

// ManagedCartridge represents a cartridge managed by the CartridgeHost.
type ManagedCartridge struct {
	path        string
	cmd         *exec.Cmd
	writerCh    chan *Frame
	manifest    []byte
	limits      Limits
	caps        []string
	knownCaps   []string
	running     bool
	helloFailed bool
}

// CartridgeHost manages N cartridge binaries with cap-based routing.
//
// Cartridges are either registered (for on-demand spawning) or attached
// (pre-connected). REQ frames from the relay are routed to the correct
// cartridge by cap URN. Continuation frames (STREAM_START, CHUNK,
// STREAM_END, END) are routed by request ID.
type CartridgeHost struct {
	cartridges []*ManagedCartridge
	capTable   []capTableEntry

	// Routing tables — mirror of the Rust `CartridgeHostRuntime`
	// design (capdag/src/bifaci/host_runtime.rs). Two independent
	// maps so self-loop peer requests (where the requesting and
	// answering cartridge are behind the same relay connection)
	// can be routed correctly:
	//
	//   outgoingRids[rid]            = cartridge that SENT the peer REQ.
	//                                  Keyed by RID alone — the relay
	//                                  hasn't assigned an XID for
	//                                  cartridge-initiated requests.
	//                                  Used to deliver the peer
	//                                  RESPONSE (frames coming back
	//                                  from the relay) to the
	//                                  requester.
	//
	//   incomingRxids[(xid, rid)]    = cartridge that RECEIVED a
	//                                  request from the relay.
	//                                  Keyed by (XID, RID) because
	//                                  for self-loop peers the same
	//                                  RID also exists in
	//                                  outgoingRids; the XID from
	//                                  the RelaySwitch disambiguates.
	//                                  Used to deliver the request
	//                                  BODY (continuation frames
	//                                  from the relay) to the
	//                                  handler.
	//
	// Phase tracking: when the request body END arrives from the
	// relay, the `incomingRxids` entry is removed — subsequent
	// frames with the same (XID, RID) fall through to
	// `outgoingRids` and are routed as peer responses. Frame
	// ordering on a single socket guarantees END is last for the
	// body phase, so the transition is unambiguous.
	outgoingRids  map[string]int    // rid string → cartridge_idx
	incomingRxids map[rxidKey]int   // (xid, rid) → cartridge_idx

	capabilities []byte
	eventCh      chan cartridgeEvent
	mu           sync.Mutex
}

// NewCartridgeHost creates a new multi-cartridge host.
func NewCartridgeHost() *CartridgeHost {
	return &CartridgeHost{
		outgoingRids:  make(map[string]int),
		incomingRxids: make(map[rxidKey]int),
		eventCh:       make(chan cartridgeEvent, 256),
	}
}

// RegisterCartridge registers a cartridge binary for on-demand spawning.
// The cartridge is not spawned until a REQ arrives for one of its known caps.
func (h *CartridgeHost) RegisterCartridge(path string, knownCaps []string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	cartridgeIdx := len(h.cartridges)
	h.cartridges = append(h.cartridges, &ManagedCartridge{
		path:      path,
		knownCaps: knownCaps,
		running:   false,
		limits:    DefaultLimits(),
	})

	for _, cap := range knownCaps {
		h.capTable = append(h.capTable, capTableEntry{capUrn: cap, cartridgeIdx: cartridgeIdx})
	}
}

// AttachCartridge attaches a pre-connected cartridge (already running).
// Performs HELLO handshake immediately and returns the cartridge index.
func (h *CartridgeHost) AttachCartridge(cartridgeRead io.Reader, cartridgeWrite io.Writer) (int, error) {
	reader := NewFrameReader(cartridgeRead)
	writer := NewFrameWriter(cartridgeWrite)

	manifest, limits, err := HandshakeInitiate(reader, writer)
	if err != nil {
		return -1, fmt.Errorf("handshake failed: %w", err)
	}

	reader.SetLimits(limits)
	writer.SetLimits(limits)

	caps, err := parseCapsFromManifest(manifest)
	if err != nil {
		return -1, fmt.Errorf("failed to parse manifest: %w", err)
	}

	h.mu.Lock()
	cartridgeIdx := len(h.cartridges)

	writerCh := make(chan *Frame, 64)
	cartridge := &ManagedCartridge{
		writerCh: writerCh,
		manifest: manifest,
		limits:   limits,
		caps:     caps,
		running:  true,
	}
	h.cartridges = append(h.cartridges, cartridge)

	for _, cap := range caps {
		h.capTable = append(h.capTable, capTableEntry{capUrn: cap, cartridgeIdx: cartridgeIdx})
	}
	h.rebuildCapabilities()
	h.mu.Unlock()

	go h.writerLoop(writer, writerCh)
	go h.readerLoop(cartridgeIdx, reader)

	return cartridgeIdx, nil
}

// Capabilities returns the aggregate capabilities of all running cartridges as JSON.
func (h *CartridgeHost) Capabilities() []byte {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.capabilities
}

// FindCartridgeForCap finds the cartridge index that can handle a given cap URN.
// Returns (cartridgeIdx, true) if found, (-1, false) if not.
func (h *CartridgeHost) FindCartridgeForCap(capUrn string) (int, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.findCartridgeForCapLocked(capUrn)
}

func (h *CartridgeHost) findCartridgeForCapLocked(capUrn string) (int, bool) {
	// Exact match first
	for _, entry := range h.capTable {
		if entry.capUrn == capUrn {
			return entry.cartridgeIdx, true
		}
	}

	// URN-level matching: use is_dispatchable (provider can handle request)
	requestUrn, err := urn.NewCapUrnFromString(capUrn)
	if err != nil {
		return -1, false
	}

	requestSpecificity := requestUrn.Specificity()

	type matchEntry struct {
		cartridgeIdx   int
		signedDistance int
	}
	var matches []matchEntry

	for _, entry := range h.capTable {
		registeredUrn, err := urn.NewCapUrnFromString(entry.capUrn)
		if err != nil {
			continue
		}
		// Use is_dispatchable: can this provider handle this request?
		if registeredUrn.IsDispatchable(requestUrn) {
			specificity := registeredUrn.Specificity()
			signedDistance := specificity - requestSpecificity
			matches = append(matches, matchEntry{entry.cartridgeIdx, signedDistance})
		}
	}

	if len(matches) == 0 {
		return -1, false
	}

	// Rank: non-negative distance (refinement/exact) before negative (fallback),
	// then by smallest absolute distance
	sort.SliceStable(matches, func(i, j int) bool {
		iGroup := 0
		if matches[i].signedDistance < 0 {
			iGroup = 1
		}
		jGroup := 0
		if matches[j].signedDistance < 0 {
			jGroup = 1
		}
		if iGroup != jGroup {
			return iGroup < jGroup
		}
		iAbs := matches[i].signedDistance
		if iAbs < 0 {
			iAbs = -iAbs
		}
		jAbs := matches[j].signedDistance
		if jAbs < 0 {
			jAbs = -jAbs
		}
		return iAbs < jAbs
	})

	return matches[0].cartridgeIdx, true
}

// Run runs the main event loop, reading from relay and cartridges.
// Blocks until relay closes or a fatal error occurs.
func (h *CartridgeHost) Run(relayRead io.Reader, relayWrite io.Writer, resourceFn func() []byte) error {
	relayReader := NewFrameReader(relayRead)
	relayWriter := NewFrameWriter(relayWrite)

	relayCh := make(chan *Frame, 64)
	relayDone := make(chan error, 1)
	go func() {
		for {
			frame, err := relayReader.ReadFrame()
			if err != nil {
				if err == io.EOF {
					relayDone <- nil
				} else {
					relayDone <- err
				}
				close(relayCh)
				return
			}
			relayCh <- frame
		}
	}()

	for {
		select {
		case frame, ok := <-relayCh:
			if !ok {
				err := <-relayDone
				h.killAllCartridges()
				return err
			}
			if err := h.handleRelayFrame(frame, relayWriter); err != nil {
				h.killAllCartridges()
				return err
			}

		case event := <-h.eventCh:
			if event.isDeath {
				h.handleCartridgeDeath(event.cartridgeIdx, relayWriter)
			} else if event.frame != nil {
				h.handleCartridgeFrame(event.cartridgeIdx, event.frame, relayWriter)
			}
		}
	}
}

// handleRelayFrame routes an incoming frame from the relay to the
// correct cartridge. Mirrors the Rust `handle_relay_frame` design
// in capdag/src/bifaci/host_runtime.rs.
//
// PATH B (REQ from relay): cap dispatch picks a cartridge, the
// (XID, RID) → cartridge_idx mapping is recorded in
// `incomingRxids`, and the frame is forwarded to the cartridge
// (still carrying the XID).
//
// PATH C (continuation frames from relay): route by checking
// `incomingRxids[(xid, rid)]` first (request body phase) and
// falling back to `outgoingRids[rid]` (peer response phase). For
// self-loop peer requests the same RID exists in both maps; the
// XID disambiguates because the body's END (which removes
// `incomingRxids[(xid, rid)]`) always precedes the peer response's
// frames on a single ordered relay socket.
func (h *CartridgeHost) handleRelayFrame(frame *Frame, relayWriter *FrameWriter) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	switch frame.FrameType {
	case FrameTypeReq:
		// All frames from the relay carry an XID assigned by the
		// RelaySwitch. Cartridge-bound REQs without one are a
		// protocol error from upstream.
		if frame.RoutingId == nil {
			return fmt.Errorf("REQ from relay missing XID — RelaySwitch must stamp routing_id")
		}
		xid := *frame.RoutingId

		capUrn := ""
		if frame.Cap != nil {
			capUrn = *frame.Cap
		}

		cartridgeIdx, found := h.findCartridgeForCapLocked(capUrn)
		if !found {
			errFrame := NewErr(frame.Id, "NO_HANDLER", fmt.Sprintf("no cartridge handles cap: %s", capUrn))
			errFrame.RoutingId = frame.RoutingId
			relayWriter.WriteFrame(errFrame)
			return nil
		}

		cartridge := h.cartridges[cartridgeIdx]
		if !cartridge.running {
			if cartridge.helloFailed {
				errFrame := NewErr(frame.Id, "SPAWN_FAILED", "cartridge previously failed to start")
				errFrame.RoutingId = frame.RoutingId
				relayWriter.WriteFrame(errFrame)
				return nil
			}
			if err := h.spawnCartridgeLocked(cartridgeIdx); err != nil {
				errFrame := NewErr(frame.Id, "SPAWN_FAILED", err.Error())
				errFrame.RoutingId = frame.RoutingId
				relayWriter.WriteFrame(errFrame)
				return nil
			}
		}

		// Record (XID, RID) → cartridge for routing the request
		// body's continuation frames. The cartridge receives the
		// REQ with XID intact so its response frames carry the
		// same XID back to the relay.
		h.incomingRxids[makeRxidKey(xid, frame.Id)] = cartridgeIdx
		h.sendToCartridge(cartridgeIdx, frame)

	case FrameTypeStreamStart, FrameTypeChunk, FrameTypeStreamEnd, FrameTypeEnd, FrameTypeErr:
		// Continuation frame from the relay. Two possibilities:
		//   1. Body phase — `incomingRxids[(xid, rid)]` says which
		//      cartridge is handling the original request. END
		//      here marks the end of the body; we drop the entry
		//      so a self-loop peer response (same RID) can fall
		//      through to outgoingRids next.
		//   2. Response phase — `outgoingRids[rid]` says which
		//      cartridge sent the peer REQ; the relay is now
		//      delivering the response back. END/ERR here marks
		//      the end of the response and drops the entry.
		//
		// MUST have XID. Frames from the relay without XID are a
		// protocol violation upstream.
		if frame.RoutingId == nil {
			return fmt.Errorf("%v from relay missing XID — all frames from relay must have XID", frame.FrameType)
		}
		xid := *frame.RoutingId
		key := makeRxidKey(xid, frame.Id)

		var (
			cartridgeIdx       int
			routedViaIncoming  bool
			haveRoute          bool
		)
		if idx, ok := h.incomingRxids[key]; ok {
			cartridgeIdx = idx
			routedViaIncoming = true
			haveRoute = true
		} else if idx, ok := h.outgoingRids[frame.Id.ToString()]; ok {
			cartridgeIdx = idx
			routedViaIncoming = false
			haveRoute = true
		}
		if !haveRoute {
			// No routing — the request was already torn down (e.g.
			// after cartridge death). Drop silently rather than
			// resurrecting state.
			return nil
		}

		h.sendToCartridge(cartridgeIdx, frame)

		isTerminal := frame.FrameType == FrameTypeEnd || frame.FrameType == FrameTypeErr
		if isTerminal {
			if routedViaIncoming {
				// Body phase done. The cartridge's response phase
				// is independent and tracked via the cartridge's
				// outbound frames carrying the same XID.
				delete(h.incomingRxids, key)
			} else {
				// Peer response phase done. Drop the requester's
				// entry so the next REQ with the same RID
				// (extremely unlikely with UUIDs but possible
				// with deterministic id allocators) starts fresh.
				delete(h.outgoingRids, frame.Id.ToString())
			}
		}

	case FrameTypeHeartbeat:
		// Engine-level heartbeat — not forwarded to cartridges.
		return nil

	case FrameTypeHello:
		return fmt.Errorf("unexpected HELLO from relay")

	case FrameTypeRelayNotify, FrameTypeRelayState:
		return fmt.Errorf("relay frame %v reached cartridge host", frame.FrameType)

	case FrameTypeLog:
		// LOG frames from peer responses — route to the cartridge
		// that made the peer request, identified by
		// `outgoingRids[rid]`. Mirrors Rust handling.
		if idx, ok := h.outgoingRids[frame.Id.ToString()]; ok {
			h.sendToCartridge(idx, frame)
		}
		return nil
	}

	return nil
}

// handleCartridgeFrame processes a frame from a cartridge. Mirrors the
// Rust `handle_cartridge_frame` design in capdag/src/bifaci/host_runtime.rs.
//
// PATH A (REQ from cartridge — peer invoke): MUST NOT carry an XID
// (cartridges never assign XIDs; the RelaySwitch does). Recorded in
// `outgoingRids[rid]` so the eventual peer response can be routed
// back to the requesting cartridge. Forwarded as-is.
//
// PATH A (continuation frames from cartridge): forwarded as-is. No
// routing-table cleanup happens here — `incomingRxids` is cleared
// only when the request BODY's END arrives from the relay (in
// `handleRelayFrame`), because cartridge response END and relay
// body END race independently.
func (h *CartridgeHost) handleCartridgeFrame(cartridgeIdx int, frame *Frame, relayWriter *FrameWriter) {
	h.mu.Lock()
	defer h.mu.Unlock()

	switch frame.FrameType {
	case FrameTypeHeartbeat:
		// Respond to cartridge heartbeat locally — don't forward.
		response := NewHeartbeat(frame.Id)
		h.sendToCartridge(cartridgeIdx, response)

	case FrameTypeHello:
		// HELLO post-handshake — protocol violation, ignore.
		return

	case FrameTypeRelayNotify, FrameTypeRelayState:
		// Cartridges must never send relay frames.
		return

	case FrameTypeReq:
		// PATH A: peer invoke. Must not carry XID.
		if frame.RoutingId != nil {
			return
		}
		h.outgoingRids[frame.Id.ToString()] = cartridgeIdx
		relayWriter.WriteFrame(frame)

	default:
		// Continuation frames (StreamStart/Chunk/StreamEnd/End/Err/Log).
		// Forward as-is, with whatever XID the cartridge stamped (it
		// echoes back the XID it received on the inbound REQ).
		relayWriter.WriteFrame(frame)
	}
}

// handleCartridgeDeath processes a cartridge death event.
func (h *CartridgeHost) handleCartridgeDeath(cartridgeIdx int, relayWriter *FrameWriter) {
	h.mu.Lock()
	defer h.mu.Unlock()

	cartridge := h.cartridges[cartridgeIdx]
	cartridge.running = false

	if cartridge.writerCh != nil {
		close(cartridge.writerCh)
		cartridge.writerCh = nil
	}

	if cartridge.cmd != nil && cartridge.cmd.Process != nil {
		cartridge.cmd.Process.Kill()
		cartridge.cmd = nil
	}

	// Send ERR for all pending requests this cartridge was involved
	// in — both the request bodies it was handling (incomingRxids)
	// and the peer requests it had outstanding (outgoingRids).
	// Mirrors Rust's heartbeat-timeout cleanup in
	// capdag/src/bifaci/host_runtime.rs:1950.
	type incomingFailure struct {
		key rxidKey
		xid MessageId
		rid MessageId
	}
	var failedIncoming []incomingFailure
	for key, idx := range h.incomingRxids {
		if idx != cartridgeIdx {
			continue
		}
		xid, errX := NewMessageIdFromUuidString(key.xid)
		rid, errR := NewMessageIdFromUuidString(key.rid)
		if errX != nil || errR != nil {
			// Non-UUID id (uint variant) — fall back via numeric parse not
			// implemented here; skip rather than fabricate. This path is
			// effectively unreachable in practice (relay always issues UUIDs).
			delete(h.incomingRxids, key)
			continue
		}
		failedIncoming = append(failedIncoming, incomingFailure{key: key, xid: xid, rid: rid})
	}
	for _, f := range failedIncoming {
		errFrame := NewErr(f.rid, "CARTRIDGE_DIED", fmt.Sprintf("cartridge %d died", cartridgeIdx))
		xid := f.xid
		errFrame.RoutingId = &xid
		relayWriter.WriteFrame(errFrame)
		delete(h.incomingRxids, f.key)
	}

	type outgoingFailure struct {
		key string
		rid MessageId
	}
	var failedOutgoing []outgoingFailure
	for key, idx := range h.outgoingRids {
		if idx != cartridgeIdx {
			continue
		}
		rid, err := NewMessageIdFromUuidString(key)
		if err != nil {
			delete(h.outgoingRids, key)
			continue
		}
		failedOutgoing = append(failedOutgoing, outgoingFailure{key: key, rid: rid})
	}
	for _, f := range failedOutgoing {
		errFrame := NewErr(f.rid, "CARTRIDGE_DIED", fmt.Sprintf("cartridge %d died", cartridgeIdx))
		relayWriter.WriteFrame(errFrame)
		delete(h.outgoingRids, f.key)
	}

	h.updateCapTable()
	h.rebuildCapabilities()
}

// sendToCartridge sends a frame to a cartridge via its writer channel.
func (h *CartridgeHost) sendToCartridge(cartridgeIdx int, frame *Frame) {
	cartridge := h.cartridges[cartridgeIdx]
	if cartridge.writerCh != nil {
		select {
		case cartridge.writerCh <- frame:
		default:
			// Channel full — cartridge probably dead, frame dropped
		}
	}
}

// writerLoop reads frames from the channel and writes them to the cartridge.
func (h *CartridgeHost) writerLoop(writer *FrameWriter, ch chan *Frame) {
	for frame := range ch {
		if err := writer.WriteFrame(frame); err != nil {
			return
		}
	}
}

// readerLoop reads frames from a cartridge and sends events to the event channel.
func (h *CartridgeHost) readerLoop(cartridgeIdx int, reader *FrameReader) {
	for {
		frame, err := reader.ReadFrame()
		if err != nil {
			h.eventCh <- cartridgeEvent{cartridgeIdx: cartridgeIdx, isDeath: true}
			return
		}
		h.eventCh <- cartridgeEvent{cartridgeIdx: cartridgeIdx, frame: frame}
	}
}

// spawnCartridgeLocked spawns a registered cartridge process (caller must hold mu).
func (h *CartridgeHost) spawnCartridgeLocked(cartridgeIdx int) error {
	cartridge := h.cartridges[cartridgeIdx]
	if cartridge.path == "" {
		cartridge.helloFailed = true
		return fmt.Errorf("cartridge has no path")
	}

	cmd := exec.Command(cartridge.path)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		cartridge.helloFailed = true
		return fmt.Errorf("failed to create stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cartridge.helloFailed = true
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		cartridge.helloFailed = true
		return fmt.Errorf("failed to start cartridge: %w", err)
	}
	cartridge.cmd = cmd

	reader := NewFrameReader(stdout)
	writer := NewFrameWriter(stdin)

	manifest, limits, err := HandshakeInitiate(reader, writer)
	if err != nil {
		cartridge.helloFailed = true
		cmd.Process.Kill()
		return fmt.Errorf("handshake failed: %w", err)
	}

	reader.SetLimits(limits)
	writer.SetLimits(limits)

	caps, parseErr := parseCapsFromManifest(manifest)
	if parseErr != nil {
		cartridge.helloFailed = true
		cmd.Process.Kill()
		return fmt.Errorf("failed to parse manifest: %w", parseErr)
	}

	cartridge.manifest = manifest
	cartridge.limits = limits
	cartridge.caps = caps
	cartridge.running = true

	writerCh := make(chan *Frame, 64)
	cartridge.writerCh = writerCh

	h.updateCapTable()
	h.rebuildCapabilities()

	go h.writerLoop(writer, writerCh)
	go h.readerLoop(cartridgeIdx, reader)

	return nil
}

// updateCapTable rebuilds the cap table from all cartridges.
func (h *CartridgeHost) updateCapTable() {
	h.capTable = nil
	for idx, cartridge := range h.cartridges {
		if cartridge.helloFailed {
			continue
		}
		caps := cartridge.knownCaps
		if cartridge.running && len(cartridge.caps) > 0 {
			caps = cartridge.caps
		}
		for _, cap := range caps {
			h.capTable = append(h.capTable, capTableEntry{capUrn: cap, cartridgeIdx: idx})
		}
	}
}

// rebuildCapabilities rebuilds the aggregate capabilities JSON.
func (h *CartridgeHost) rebuildCapabilities() {
	var allCaps []string
	for _, cartridge := range h.cartridges {
		if cartridge.running {
			allCaps = append(allCaps, cartridge.caps...)
		}
	}

	if len(allCaps) == 0 {
		h.capabilities = nil
		return
	}

	payload := RelayNotifyCapabilitiesPayload{
		Caps:                allCaps,
		InstalledCartridges: []InstalledCartridgeIdentity{},
	}
	capsJSON, err := json.Marshal(payload)
	if err != nil {
		h.capabilities = nil
		return
	}
	h.capabilities = capsJSON
}

// killAllCartridges stops all managed cartridges.
func (h *CartridgeHost) killAllCartridges() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, cartridge := range h.cartridges {
		if cartridge.writerCh != nil {
			close(cartridge.writerCh)
			cartridge.writerCh = nil
		}
		if cartridge.cmd != nil && cartridge.cmd.Process != nil {
			cartridge.cmd.Process.Kill()
		}
		cartridge.running = false
	}
}

// parseCapsFromManifest parses cap URNs from a JSON manifest.
// Expected format: {"cap_groups": [{"caps": [{"urn": "cap:op=test", ...}, ...]}]}
func parseCapsFromManifest(manifest []byte) ([]string, error) {
	if len(manifest) == 0 {
		return nil, nil
	}

	var parsed struct {
		CapGroups []struct {
			Caps []struct {
				Urn string `json:"urn"`
			} `json:"caps"`
		} `json:"cap_groups"`
	}

	if err := json.Unmarshal(manifest, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse manifest JSON: %w", err)
	}

	if len(parsed.CapGroups) == 0 {
		return nil, fmt.Errorf("manifest missing required cap_groups array")
	}

	var caps []string
	for _, group := range parsed.CapGroups {
		for _, cap := range group.Caps {
			if cap.Urn != "" {
				caps = append(caps, cap.Urn)
			}
		}
	}

	return caps, nil
}
