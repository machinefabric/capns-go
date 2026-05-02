package bifaci

// CartridgeHostObserver receives lifecycle notifications from a CartridgeHost.
// All methods are called with the host mutex NOT held; implementations must
// not call back into the CartridgeHost.
type CartridgeHostObserver interface {
	// CartridgeSpawned is called when a cartridge process starts and
	// completes its HELLO handshake successfully.
	CartridgeSpawned(cartridgeIndex int, pid *uint32, name string, caps []string)

	// CartridgeDied is called when a cartridge process exits or its
	// connection is lost.
	CartridgeDied(cartridgeIndex int, pid *uint32, name string)
}
