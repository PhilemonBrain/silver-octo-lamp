package main

import (
	"log"

	"github.com/PhilemonBrain/d-file-storage/p2p"
)

// func main() {
// 	tcpOpts := p2p.TCPTransportOptions{
// 		ListenAddress: ":3000",
// 		ShakeHands:    p2p.NOPHandshakeFunc,
// 		Decoder:       p2p.DefaultDecoder{},
// 	}
// 	tr := p2p.NewTCPTransport(tcpOpts)

// 	go func() {
// 		ch := tr.Consume()
// 		for {
// 			msg := <-ch
// 			fmt.Printf("%+v \n", msg)
// 		}
// 	}()

// 	if err := tr.ListenAndAccept(); err != nil {
// 		log.Fatal("Somethnig jsut happepede")
// 	}
// 	fmt.Println("We Live!")

// 	select {}

// }

func main() {
	transportOpts := p2p.TCPTransportOptions{
		ListenAddress: ":3000",
		ShakeHands:    p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	transport := p2p.NewTCPTransport(transportOpts)
	fileServerOpts := FileServerOpts{
		ListenAddr:        ":3000",
		StorageRoot:       ":3000_network",
		pathTransformFunc: CASPathTransformFunc,
		Transport:         transport,
	}
	server := NewFileServer(fileServerOpts)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
