package main

import (
	"fmt"
	"log"

	"github.com/PhilemonBrain/d-file-storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    p2p.NOPHandshake,
		// Decoder: ///,

	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal("Somethnig jsut happepede")
	}
	fmt.Println("We Live!")

	select {}

}
