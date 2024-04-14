package p2p

// HandshakeFunc
type HanshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
