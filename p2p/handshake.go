package p2p

type HandShakerFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
