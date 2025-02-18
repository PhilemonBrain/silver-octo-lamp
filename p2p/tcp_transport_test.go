package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := &TCPTransportOptions{
		ListenAddress: ":3000",
		Decoder:       DefaultDecoder{},
		ShakeHands:    NOPHandshakeFunc,
	}
	tcpTransport := NewTCPTransport(*opts)
	assert.Equal(t, opts.ListenAddress, tcpTransport.ListenAddress)
}
