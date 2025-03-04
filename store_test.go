package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		pathTransformFunc: DefaultPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some bytes"))
	if err := s.writeStream("myspecialP", data); err != nil {
		t.Error(err)
	}
}
