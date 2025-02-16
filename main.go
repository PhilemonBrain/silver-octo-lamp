package main

import (
	"fmt"
	"log"

	"github.com/PhilemonBrain/d-file-storage/p2p"
)

func main() {

	tr := p2p.NewTCPTransport(":4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal("Somethnig jsut happepede")
	}
	fmt.Println("We Live!")

	select {}

}
