package p2p

import (
	// "bytes"
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPPeer represents a node in a TCP connection
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// outbound is True when we dial a connection,
	// but false when we accept a connection
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOptions struct {
	ListenAddress string
	ShakeHands    HandShakerFunc
	Decoder       Decoder
}

type TCPTransport struct {
	listener net.Listener
	TCPTransportOptions

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		log.Fatal("Something wrong at TCP transport level")
		return err
	}

	go t.startAcceptLoop()

	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("tcp accept error %s\n", err)
		}

		fmt.Printf("new incoming connection %+v\n", conn)

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	// peer := NewTCPPeer(conn, true)

	if err := t.ShakeHands(conn); err != nil {
		fmt.Println("doingn somehting wrong")
		conn.Close()
		return
	}

	msg := &Message{}

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error %s/n", err)
			continue
		}
		msg.From = conn.RemoteAddr()
		fmt.Printf(" message :%+v \n", msg)
	}
}
