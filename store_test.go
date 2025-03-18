package main

import (
	"bytes"
	"fmt"
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
	s := createNewStore()
	defer tearDown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("randkomTestingKeviu34f09j490fj94f3y%d", i)

		data := []byte("this is the actual contents to be written to file")
		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("expected to have key %s", key)
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

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		// after delete we need to retest that there is no key
		if ok := s.Has(key); ok {
			t.Errorf("expected not to have key %s", key)
		}
	}

}

func TestDelete(t *testing.T) {
	s := createNewStore()
	defer tearDown(t, s)

	key := "randkomTestingKeviu34f09j490fj94f3y"
	data := []byte("this is the actual contents to be written to file")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(key)
	}
}

func createNewStore() *Store {
	opts := StoreOpts{
		pathTransformFunc: CASPathTransportFunc,
	}
	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
