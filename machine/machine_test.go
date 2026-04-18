package machine

import (
	"testing"

	"github.com/machinefabric/capdag-go/cap"
	"github.com/machinefabric/capdag-go/planner"
	"github.com/machinefabric/capdag-go/urn"
)

// ===================================================================
// Test fixtures
// ===================================================================

// buildCap constructs a *cap.Cap with a single-stdin-arg per entry in
// argMediaUrns. Slot identity == stdin URN for each arg.
func buildCap(capUrnStr, title string, argMediaUrns []string, outputMediaUrn string) *cap.Cap {
	capUrnParsed, err := urn.NewCapUrnFromString(capUrnStr)
	if err != nil {
		panic("test fixture: invalid cap URN " + capUrnStr + ": " + err.Error())
	}
	args := make([]cap.CapArg, len(argMediaUrns))
	for i, mu := range argMediaUrns {
		stdinVal := mu
		args[i] = cap.NewCapArg(mu, true, []cap.ArgSource{{Stdin: &stdinVal}})
	}
	outMedia := outputMediaUrn
	output := cap.NewCapOutput(outMedia, "output of "+title)
	return &cap.Cap{
		Urn:     capUrnParsed,
		Title:   title,
		Command: "test-fixture://" + title,
		Args:    args,
		Output:  output,
	}
}

// registryWith builds a test CapRegistry pre-populated with the given caps.
func registryWith(caps []*cap.Cap) *cap.CapRegistry {
	r := cap.NewCapRegistryForTest()
	r.AddCapsToCache(caps)
	return r
}

// mediaUrn parses a media URN string; panics on failure.
func mediaUrn(s string) *urn.MediaUrn {
	m, err := urn.NewMediaUrnFromString(s)
	if err != nil {
		panic("test fixture: invalid media URN " + s + ": " + err.Error())
	}
	return m
}

// capUrnVal parses a cap URN string; panics on failure.
func capUrnVal(s string) *urn.CapUrn {
	c, err := urn.NewCapUrnFromString(s)
	if err != nil {
		panic("test fixture: invalid cap URN " + s + ": " + err.Error())
	}
	return c
}

// capStep builds a StepTypeCap StrandStep.
func capStep(capUrnStr, title, from, to string) *planner.StrandStep {
	return &planner.StrandStep{
		StepType:  planner.StepTypeCap,
		FromSpec:  mediaUrn(from),
		ToSpec:    mediaUrn(to),
		CapUrnVal: capUrnVal(capUrnStr),
		StepTitle: title,
	}
}

// strandFromSteps wraps steps into a Strand.
func strandFromSteps(steps []*planner.StrandStep, description string) *planner.Strand {
	totalSteps := len(steps)
	capStepCount := 0
	for _, s := range steps {
		if s.StepType == planner.StepTypeCap {
			capStepCount++
		}
	}
	sourceSpec := steps[0].FromSpec
	targetSpec := steps[len(steps)-1].ToSpec
	return &planner.Strand{
		Steps:        steps,
		SourceSpec:   sourceSpec,
		TargetSpec:   targetSpec,
		TotalSteps:   totalSteps,
		CapStepCount: capStepCount,
		Description:  description,
	}
}

// ===================================================================
// Cap definitions used across tests
// ===================================================================

func extractCapDef() *cap.Cap {
	return buildCap(
		`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
		"extract",
		[]string{"media:pdf"},
		`media:txt;textable`,
	)
}

func embedCapDef() *cap.Cap {
	return buildCap(
		`cap:in=media:textable;op=embed;out="media:vec;record"`,
		"embed",
		[]string{"media:textable"},
		`media:vec;record`,
	)
}

func pdfToTxtStrand() *planner.Strand {
	return strandFromSteps(
		[]*planner.StrandStep{
			capStep(
				`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
				"extract",
				"media:pdf",
				`media:txt;textable`,
			),
		},
		"pdf to txt",
	)
}

func txtToVecStrand() *planner.Strand {
	return strandFromSteps(
		[]*planner.StrandStep{
			capStep(
				`cap:in=media:textable;op=embed;out="media:vec;record"`,
				"embed",
				`media:txt;textable`,
				`media:vec;record`,
			),
		},
		"txt to vec",
	)
}

// ===================================================================
// FromStrand tests
// ===================================================================

// TestFromStrandProducesSingleStrandMachine verifies that a single-step
// planner strand yields a Machine with exactly one MachineStrand and one edge.
// TEST1155: Building a machine from one strand produces one strand with one resolved edge.
func Test1155_FromStrandProducesSingleStrandMachine(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	m, err := FromStrand(pdfToTxtStrand(), registry)
	if err != nil {
		t.Fatalf("FromStrand failed: %s", err)
	}
	if m.StrandCount() != 1 {
		t.Fatalf("expected 1 strand, got %d", m.StrandCount())
	}
	if len(m.Strands()[0].Edges()) != 1 {
		t.Fatalf("expected 1 edge in strand, got %d", len(m.Strands()[0].Edges()))
	}
}

// TestFromStrandsKeepStrandsDisjoint verifies that FromStrands does NOT join
// two strands even when their URNs are type-compatible at runtime.
// TEST1156: Building from multiple strands keeps them disjoint and preserves input strand order.
func Test1156_FromStrandsKeepStrandsDisjoint(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef(), embedCapDef()})
	m, err := FromStrands([]*planner.Strand{pdfToTxtStrand(), txtToVecStrand()}, registry)
	if err != nil {
		t.Fatalf("FromStrands failed: %s", err)
	}
	if m.StrandCount() != 2 {
		t.Fatalf("FromStrands must keep input strands as disjoint MachineStrands; got %d", m.StrandCount())
	}
	if len(m.Strands()[0].Edges()) != 1 {
		t.Fatalf("strand 0: expected 1 edge, got %d", len(m.Strands()[0].Edges()))
	}
	if len(m.Strands()[1].Edges()) != 1 {
		t.Fatalf("strand 1: expected 1 edge, got %d", len(m.Strands()[1].Edges()))
	}
	// Strand order must match input order.
	if !containsStr(m.Strands()[0].Edges()[0].CapUrn.String(), "op=extract") {
		t.Errorf("strand 0 should use extract cap, got %s", m.Strands()[0].Edges()[0].CapUrn)
	}
	if !containsStr(m.Strands()[1].Edges()[0].CapUrn.String(), "op=embed") {
		t.Errorf("strand 1 should use embed cap, got %s", m.Strands()[1].Edges()[0].CapUrn)
	}
}

// TestFromStrandsEmptyInputFailsHard verifies that passing an empty slice to
// FromStrands returns a NoCapabilitySteps error.
// TEST1157: Building from zero strands fails with NoCapabilitySteps.
func Test1157_FromStrandsEmptyInputFailsHard(t *testing.T) {
	registry := registryWith([]*cap.Cap{})
	_, err := FromStrands([]*planner.Strand{}, registry)
	if err == nil {
		t.Fatal("expected error for empty strands, got nil")
	}
	if err.Kind != ErrAbstractionNoCapabilitySteps {
		t.Errorf("expected ErrAbstractionNoCapabilitySteps, got %v", err.Kind)
	}
}

// TestMachineIsEquivalentIsStrictPositional verifies that swapping strand
// order breaks IsEquivalent — strand declaration order is part of identity.
// TEST1158: Machine equivalence is strict about strand order and rejects reordered strands.
func Test1158_MachineIsEquivalentIsStrictPositional(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef(), embedCapDef()})
	forward, err := FromStrands([]*planner.Strand{pdfToTxtStrand(), txtToVecStrand()}, registry)
	if err != nil {
		t.Fatal(err)
	}
	reversed, err := FromStrands([]*planner.Strand{txtToVecStrand(), pdfToTxtStrand()}, registry)
	if err != nil {
		t.Fatal(err)
	}
	if forward.IsEquivalent(reversed) {
		t.Error("swapping strand order must break strict equivalence")
	}
	if !forward.IsEquivalent(forward) {
		t.Error("a machine must be equivalent to itself")
	}
	if !reversed.IsEquivalent(reversed) {
		t.Error("a machine must be equivalent to itself")
	}
}

// TestMachineStrandIsEquivalentWalksNodeBijection verifies that two
// MachineStrands built from the same planner strand are equivalent.
// TEST1159: MachineStrand equivalence accepts two separately built but structurally identical strands.
func Test1159_MachineStrandIsEquivalentWalksNodeBijection(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	m1, err1 := FromStrand(pdfToTxtStrand(), registry)
	m2, err2 := FromStrand(pdfToTxtStrand(), registry)
	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected error: %v %v", err1, err2)
	}
	if !m1.Strands()[0].IsEquivalent(m2.Strands()[0]) {
		t.Error("two MachineStrands built from the same planner strand must be equivalent")
	}
}

// ===================================================================
// Anchor computation tests
// ===================================================================

// TestInputOutputAnchors verifies that the resolver correctly identifies
// root (input) and leaf (output) nodes for a simple linear strand.
// TEST1160: Creating a MachineRun stores the canonical notation and starts in the pending state.
func Test1160_InputOutputAnchors(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	m, err := FromStrand(pdfToTxtStrand(), registry)
	if err != nil {
		t.Fatal(err)
	}
	strand := m.Strands()[0]

	if len(strand.InputAnchorIds()) != 1 {
		t.Errorf("expected 1 input anchor, got %d", len(strand.InputAnchorIds()))
	}
	if len(strand.OutputAnchorIds()) != 1 {
		t.Errorf("expected 1 output anchor, got %d", len(strand.OutputAnchorIds()))
	}

	// Input anchor URN must be media:pdf (the from_spec of the extract step).
	inputAnchors := strand.InputAnchors()
	if len(inputAnchors) != 1 || !containsStr(inputAnchors[0].String(), "pdf") {
		t.Errorf("expected input anchor to contain 'pdf', got %v", inputAnchors)
	}

	// Output anchor URN must be media:txt;textable (the to_spec).
	outputAnchors := strand.OutputAnchors()
	if len(outputAnchors) != 1 || !containsStr(outputAnchors[0].String(), "txt") {
		t.Errorf("expected output anchor to contain 'txt', got %v", outputAnchors)
	}
}

// ===================================================================
// IsLoop / ForEach tests
// ===================================================================

// TestForEachSetsIsLoop verifies that a ForEach step preceding a Cap step
// sets IsLoop=true on the resulting MachineEdge.
// TEST1169: Loop markers in notation set the resolved edge loop flag on the following cap step.
func Test1169_ForEachSetsIsLoop(t *testing.T) {
	loopCap := buildCap(
		`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
		"extract",
		[]string{"media:pdf"},
		`media:txt;textable`,
	)
	registry := registryWith([]*cap.Cap{loopCap})

	steps := []*planner.StrandStep{
		{
			StepType:  planner.StepTypeForEach,
			FromSpec:  mediaUrn("media:pdf"),
			ToSpec:    mediaUrn("media:pdf"),
			MediaSpec: mediaUrn("media:pdf"),
		},
		capStep(
			`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
			"extract",
			"media:pdf",
			`media:txt;textable`,
		),
	}
	strand := strandFromSteps(steps, "loop extract")
	m, err := FromStrand(strand, registry)
	if err != nil {
		t.Fatal(err)
	}
	edge := m.Strands()[0].Edges()[0]
	if !edge.IsLoop {
		t.Error("expected IsLoop=true on edge following ForEach step")
	}
}

// TestCollectIsElided verifies that a Collect step produces no MachineEdge —
// the resolved strand has only one edge (from the Cap step).
// TEST1170: Parsing and then serializing machine notation round-trips to the canonical form.
func Test1170_CollectIsElided(t *testing.T) {
	loopCap := buildCap(
		`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
		"extract",
		[]string{"media:pdf"},
		`media:txt;textable`,
	)
	registry := registryWith([]*cap.Cap{loopCap})

	steps := []*planner.StrandStep{
		capStep(
			`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
			"extract",
			"media:pdf",
			`media:txt;textable`,
		),
		{
			StepType:  planner.StepTypeCollect,
			FromSpec:  mediaUrn(`media:txt;textable`),
			ToSpec:    mediaUrn(`media:txt;textable`),
			MediaSpec: mediaUrn(`media:txt;textable`),
		},
	}
	strand := strandFromSteps(steps, "extract then collect")
	m, err := FromStrand(strand, registry)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Strands()[0].Edges()) != 1 {
		t.Errorf("Collect must produce no edge; expected 1 edge total, got %d", len(m.Strands()[0].Edges()))
	}
}

// ===================================================================
// Parser tests
// ===================================================================

// TestParseSingleStrandTwoCapsConnectedViaSharedNode verifies that two wirings
// sharing the node name `txt` become a single connected component (one strand).
// TEST1163: Parsing one connected strand yields a single machine strand with both caps connected by the shared node.
func Test1163_ParseSingleStrandTwoCapsConnectedViaSharedNode(t *testing.T) {
	registry := pdfExtractEmbedRegistry()
	notation := `[extract cap:in=media:pdf;op=extract;out="media:txt;textable"]` +
		`[embed cap:in=media:textable;op=embed;out="media:vec;record"]` +
		`[doc -> extract -> txt]` +
		`[txt -> embed -> vec]`

	m, parseErr := ParseMachine(notation, registry)
	if parseErr != nil {
		t.Fatalf("ParseMachine failed: %s", parseErr)
	}
	if m.StrandCount() != 1 {
		t.Fatalf("expected 1 strand (shared node 'txt' merges both wirings), got %d", m.StrandCount())
	}
	strand := m.Strands()[0]
	if len(strand.Edges()) != 2 {
		t.Fatalf("expected 2 edges in strand, got %d", len(strand.Edges()))
	}
	// The intermediate node must be the same NodeId for both edges.
	extractTarget := strand.Edges()[0].Target
	embedSource := strand.Edges()[1].Assignment[0].Source
	if extractTarget != embedSource {
		t.Errorf("intermediate node 'txt' must be the same NodeId: extract.Target=%d, embed.source=%d",
			extractTarget, embedSource)
	}
}

// TestParseTwoDisconnectedStrandsYieldsTwoMachineStrands verifies that wirings
// with no shared node names are partitioned into two separate strands.
// TEST1164: Parsing two disconnected strand definitions yields two separate machine strands.
func Test1164_ParseTwoDisconnectedStrandsYieldsTwoMachineStrands(t *testing.T) {
	convertA := buildCap(
		"cap:in=media:json;op=convert_a;out=media:csv",
		"convert_a",
		[]string{"media:json"},
		"media:csv",
	)
	convertB := buildCap(
		"cap:in=media:html;op=convert_b;out=media:txt",
		"convert_b",
		[]string{"media:html"},
		"media:txt",
	)
	registry := registryWith([]*cap.Cap{convertA, convertB})

	notation := `[ca cap:in=media:json;op=convert_a;out=media:csv]` +
		`[cb cap:in=media:html;op=convert_b;out=media:txt]` +
		`[input_a -> ca -> output_a]` +
		`[input_b -> cb -> output_b]`

	m, parseErr := ParseMachine(notation, registry)
	if parseErr != nil {
		t.Fatalf("ParseMachine failed: %s", parseErr)
	}
	if m.StrandCount() != 2 {
		t.Fatalf("two wirings sharing no nodes must produce 2 strands, got %d", m.StrandCount())
	}
	// Strand order is first-appearance order.
	if len(m.Strands()[0].Edges()) != 1 {
		t.Errorf("strand 0: expected 1 edge, got %d", len(m.Strands()[0].Edges()))
	}
	if len(m.Strands()[1].Edges()) != 1 {
		t.Errorf("strand 1: expected 1 edge, got %d", len(m.Strands()[1].Edges()))
	}
	// First strand uses convert_a, second uses convert_b.
	if !containsStr(m.Strands()[0].Edges()[0].CapUrn.String(), "convert_a") {
		t.Errorf("strand 0 should use convert_a, got %s", m.Strands()[0].Edges()[0].CapUrn)
	}
	if !containsStr(m.Strands()[1].Edges()[0].CapUrn.String(), "convert_b") {
		t.Errorf("strand 1 should use convert_b, got %s", m.Strands()[1].Edges()[0].CapUrn)
	}
}

// TestParseEmptyInputReturnsError verifies that empty/whitespace input fails
// with an ErrEmpty error.
// TEST1171: Empty machine notation is rejected as a syntax error.
func Test1171_ParseEmptyInputReturnsError(t *testing.T) {
	registry := registryWith([]*cap.Cap{})
	_, err := ParseMachine("   ", registry)
	if err == nil {
		t.Fatal("expected error for empty input")
	}
	if err.Syntax == nil || err.Syntax.Kind != ErrEmpty {
		t.Errorf("expected ErrEmpty syntax error, got %v", err)
	}
}

// TestParseHeadersWithNoWiringsReturnsNoEdgesError verifies the ErrNoEdges case.
func TestParseHeadersWithNoWiringsReturnsNoEdgesError(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	notation := `[extract cap:in=media:pdf;op=extract;out="media:txt;textable"]`
	_, err := ParseMachine(notation, registry)
	if err == nil {
		t.Fatal("expected error for headers with no wirings")
	}
	if err.Syntax == nil || err.Syntax.Kind != ErrNoEdges {
		t.Errorf("expected ErrNoEdges, got %v", err)
	}
}

// TestParseDuplicateAliasReturnsError verifies that two headers with the same
// alias name return ErrDuplicateAlias.
// TEST1166: Duplicate header aliases are reported as syntax errors.
func Test1166_ParseDuplicateAliasReturnsError(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	notation := `[extract cap:in=media:pdf;op=extract;out="media:txt;textable"]` +
		`[extract cap:in=media:pdf;op=extract;out="media:txt;textable"]` +
		`[doc -> extract -> txt]`
	_, err := ParseMachine(notation, registry)
	if err == nil {
		t.Fatal("expected error for duplicate alias")
	}
	if err.Syntax == nil || err.Syntax.Kind != ErrDuplicateAlias {
		t.Errorf("expected ErrDuplicateAlias, got %v", err)
	}
}

// TestParseUndefinedAliasReturnsError verifies that a wiring referencing an
// undefined cap alias returns ErrUndefinedAlias.
// TEST1167: Wiring that references an undefined alias is reported as a syntax error.
func Test1167_ParseUndefinedAliasReturnsError(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	notation := `[doc -> no_such_cap -> txt]`
	_, err := ParseMachine(notation, registry)
	if err == nil {
		t.Fatal("expected error for undefined alias")
	}
	if err.Syntax == nil || err.Syntax.Kind != ErrUndefinedAlias {
		t.Errorf("expected ErrUndefinedAlias, got %v", err)
	}
}

// TEST1165: Parsing fails hard when a referenced cap is missing from the registry cache.
func Test1165_ParseUnknownCapInRegistryReturnsAbstractionError(t *testing.T) {
	// Empty registry — cap won't be found during resolution.
	registry := registryWith([]*cap.Cap{})
	notation := `[ex cap:in=media:pdf;op=extract;out="media:txt;textable"]` +
		`[doc -> ex -> txt]`
	_, err := ParseMachine(notation, registry)
	if err == nil {
		t.Fatal("expected error for unknown cap in registry")
	}
	if err.Abstraction == nil || err.Abstraction.Kind != ErrAbstractionUnknownCap {
		t.Errorf("expected ErrAbstractionUnknownCap, got %v", err)
	}
}

// TestParseNodeNameCollidesWithCapAlias verifies that a node name matching a
// cap alias returns ErrNodeAliasCollision.
// TEST1168: Parsing rejects node names that collide with declared cap aliases.
func Test1168_ParseNodeNameCollidesWithCapAlias(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	// Node name 'extract' collides with cap alias 'extract'.
	notation := `[extract cap:in=media:pdf;op=extract;out="media:txt;textable"]` +
		`[extract -> extract -> txt]`
	_, err := ParseMachine(notation, registry)
	if err == nil {
		t.Fatal("expected error for node-alias collision")
	}
	if err.Syntax == nil || err.Syntax.Kind != ErrNodeAliasCollision {
		t.Errorf("expected ErrNodeAliasCollision, got %v", err)
	}
}

// ===================================================================
// Serializer tests
// ===================================================================

// TestToMachineNotationRoundTrips verifies that a machine parsed from
// notation and re-serialized produces a machine equivalent to the original.
// TEST1173: Serializing and reparsing a machine preserves strict machine equivalence.
func Test1173_ToMachineNotationRoundTrips(t *testing.T) {
	registry := pdfExtractEmbedRegistry()
	notation := `[extract cap:in=media:pdf;op=extract;out="media:txt;textable"]` +
		`[embed cap:in=media:textable;op=embed;out="media:vec;record"]` +
		`[doc -> extract -> txt]` +
		`[txt -> embed -> vec]`

	m1, parseErr := ParseMachine(notation, registry)
	if parseErr != nil {
		t.Fatalf("first parse failed: %s", parseErr)
	}

	serialized := m1.ToMachineNotation()
	if serialized == "" {
		t.Fatal("ToMachineNotation returned empty string for non-empty machine")
	}

	m2, parseErr2 := ParseMachine(serialized, registry)
	if parseErr2 != nil {
		t.Fatalf("second parse (of serialized notation) failed: %s", parseErr2)
	}

	if !m1.IsEquivalent(m2) {
		t.Errorf("round-tripped machine is not equivalent to the original.\nOriginal: %s\nSerialized: %s",
			notation, serialized)
	}
}

// TestEmptyMachineSerializesToEmpty verifies that an empty machine produces
// an empty string from ToMachineNotation.
// TEST1175: Serializing an empty machine produces an empty string.
func Test1175_EmptyMachineSerializesToEmpty(t *testing.T) {
	m := fromResolvedStrands(nil)
	if m.ToMachineNotation() != "" {
		t.Errorf("empty machine must serialize to empty string, got %q", m.ToMachineNotation())
	}
}

// TestMachineStringRepr verifies the String() representation of a machine.
// TEST1172: Serializing a two-step strand emits the expected aliases and node names.
func Test1172_MachineStringRepr(t *testing.T) {
	registry := registryWith([]*cap.Cap{extractCapDef()})
	m, err := FromStrand(pdfToTxtStrand(), registry)
	if err != nil {
		t.Fatal(err)
	}
	s := m.String()
	if !containsStr(s, "1 strands") || !containsStr(s, "1 edges") {
		t.Errorf("unexpected String() output: %q", s)
	}
}

// ===================================================================
// IsEquivalent structural corner cases
// ===================================================================

// TestStrandEquivalenceWithDifferentNodeAllocationOrders verifies that two
// equivalent strands remain equivalent even when their NodeIds were allocated
// in different positions (bijection check).
// TEST1189: Strand resolution keeps canonical anchor ordering stable across equivalent inputs.
func Test1189_StrandEquivalenceWithDifferentNodeAllocationOrders(t *testing.T) {
	// Build two machines from identical strands — node allocation order is
	// deterministic but this confirms the bijection handles it correctly.
	registry := registryWith([]*cap.Cap{extractCapDef()})
	m1, _ := FromStrand(pdfToTxtStrand(), registry)
	m2, _ := FromStrand(pdfToTxtStrand(), registry)
	if !m1.IsEquivalent(m2) {
		t.Error("identical strands must be equivalent")
	}

	// A two-step chain: extract then embed.
	twoStepCap := buildCap(
		`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
		"extract",
		[]string{"media:pdf"},
		`media:txt;textable`,
	)
	twoStepEmbed := buildCap(
		`cap:in=media:textable;op=embed;out="media:vec;record"`,
		"embed",
		[]string{"media:textable"},
		`media:vec;record`,
	)
	twoStepRegistry := registryWith([]*cap.Cap{twoStepCap, twoStepEmbed})

	twoStepStrand := strandFromSteps(
		[]*planner.StrandStep{
			capStep(`cap:in=media:pdf;op=extract;out="media:txt;textable"`, "extract", "media:pdf", `media:txt;textable`),
			capStep(`cap:in=media:textable;op=embed;out="media:vec;record"`, "embed", `media:txt;textable`, `media:vec;record`),
		},
		"extract then embed",
	)

	m3, err := FromStrand(twoStepStrand, twoStepRegistry)
	if err != nil {
		t.Fatal(err)
	}
	m4, err := FromStrand(twoStepStrand, twoStepRegistry)
	if err != nil {
		t.Fatal(err)
	}
	if !m3.IsEquivalent(m4) {
		t.Error("two-step strands built from identical input must be equivalent")
	}
}

// TestStrandNonEquivalenceDifferentCap verifies that strands with different
// cap URNs are not equivalent.
// TEST1187: Strand resolution fails when a referenced cap is not found in the registry.
func Test1187_StrandNonEquivalenceDifferentCap(t *testing.T) {
	cap1 := buildCap("cap:in=media:pdf;op=extract;out=media:txt", "extract", []string{"media:pdf"}, "media:txt")
	cap2 := buildCap("cap:in=media:pdf;op=convert;out=media:txt", "convert", []string{"media:pdf"}, "media:txt")
	reg1 := registryWith([]*cap.Cap{cap1})
	reg2 := registryWith([]*cap.Cap{cap2})

	s1, err1 := FromStrand(
		strandFromSteps([]*planner.StrandStep{
			capStep("cap:in=media:pdf;op=extract;out=media:txt", "extract", "media:pdf", "media:txt"),
		}, "s1"), reg1,
	)
	s2, err2 := FromStrand(
		strandFromSteps([]*planner.StrandStep{
			capStep("cap:in=media:pdf;op=convert;out=media:txt", "convert", "media:pdf", "media:txt"),
		}, "s2"), reg2,
	)
	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected errors: %v %v", err1, err2)
	}
	if s1.IsEquivalent(s2) {
		t.Error("strands with different cap URNs must not be equivalent")
	}
}

// ===================================================================
// Helpers
// ===================================================================

func pdfExtractEmbedRegistry() *cap.CapRegistry {
	extract := buildCap(
		`cap:in=media:pdf;op=extract;out="media:txt;textable"`,
		"extract",
		[]string{"media:pdf"},
		`media:txt;textable`,
	)
	embed := buildCap(
		`cap:in=media:textable;op=embed;out="media:vec;record"`,
		"embed",
		[]string{"media:textable"},
		`media:vec;record`,
	)
	return registryWith([]*cap.Cap{extract, embed})
}

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || findSubstr(s, substr))
}

func findSubstr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
