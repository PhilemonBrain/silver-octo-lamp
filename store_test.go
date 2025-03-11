package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "randkomTestingKeviu34f09j490fj94f3y"
	pathKey := CASPathTransportFunc(key)
	expectedOriginalKey := "bbe5751b6b689f4ad49821a96404996dd0ddf90f"
	expectedPathName := "bbe57/51b6b/689f4/ad498/21a96/40499/6dd0d/df90f"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s; want %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.FileName != expectedOriginalKey {
		t.Errorf("have %s; want %s", pathKey.FileName, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		pathTransformFunc: CASPathTransportFunc,
	}
	s := NewStore(opts)
	key := "randkomTestingKeviu34f09j490fj94f3y"

	data := []byte("this is the actual contents to be written to file")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	f, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	c, err := io.ReadAll(f)
	if err != nil {
		t.Error(err)
	}

	if string(c) != string(data) {
		t.Errorf("want %s; have %s", data, c)
	}

	s.Delete(key)

}

func TestDelete(t *testing.T) {
	storeOpts := StoreOpts{
		pathTransformFunc: CASPathTransportFunc,
	}
	s := NewStore(storeOpts)

	key := "randkomTestingKeviu34f09j490fj94f3y"
	data := []byte("this is the actual contents to be written to file")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(key)
	}
}
