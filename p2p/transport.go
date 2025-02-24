package p2p

// Peer is a repersentation of a remote node
type Peer interface {
	Close() error
}

// Transport is anything that handles a communication in a network
// eg. TCP, UDP, WebSockets
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
