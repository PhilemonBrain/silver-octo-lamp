package main

import (
	// "crypto/sha1"
	// "encoding/hex"

	"io"
	"log"
	"os"
)

// func CASPathTransformFunc(key string){
// 	hash := sha1.Sum([]byte(key))
// 	hashStr := hex.EncodeToString(hash[:])
// }

type PathTransportFunc func(string) string

type Store struct {
	StoreOpts
}

type StoreOpts struct {
	pathTransformFunc PathTransportFunc
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (store *Store) writeStream(key string, r io.Reader) error {

	pathname := store.pathTransformFunc(key)
	if err := os.Mkdir(pathname, os.ModePerm); err != nil {
		return err
	}

	filename := "randomfile"
	pathAndFilename := pathname + "/" + filename
	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("written %d bytes to disk: %s", n, pathname)

	return nil
}
