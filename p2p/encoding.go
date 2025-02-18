package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	if err != nil {
		fmt.Printf("NOPDecoder error %s\n", err)
	}
	msg.Payload = buf[:n]
	return nil
}
