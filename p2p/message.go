package p2p

import "net"

// Message holds any arbitrary data that is sent over each transport
// betweeen two nodes in a network
type RPC struct {
	From    net.Addr
	Payload []byte
}
