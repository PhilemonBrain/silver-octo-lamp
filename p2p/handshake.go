package p2p

type HandShakerFunc func(Peer) error

func NOPHandshake(Peer) error { return nil }
