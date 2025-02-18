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

// Close implements the peer interface
func (p *TCPPeer) Close() error {
	return p.conn.Close()
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

	rpcChan chan RPC

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
		rpcChan:             make(chan RPC),
	}
}

// Consume implements the transport interface
// which will return a read only channel
// for readding the incoming messages recieved from another peer
// in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcChan
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
	peer := NewTCPPeer(conn, true)

	if err := t.ShakeHands(peer); err != nil {
		fmt.Println("doingn somehting wrong")
		conn.Close()
		return
	}

	rpc := RPC{}

	for {
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error %s/n", err)
			continue
		}
		rpc.From = conn.RemoteAddr()
		t.rpcChan <- rpc
		fmt.Printf(" message :%+v \n", rpc)
	}
}
