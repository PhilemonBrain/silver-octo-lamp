package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type PathTransformFunc func(string) PathKey

// Content Adrresable storage
func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashString) / blockSize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashString[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		FileName: hashString,
	}

}

type Store struct {
	StoreOpts
}

type PathKey struct {
	PathName string
	FileName string
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.FileName)
}

func (p PathKey) RootFolder() string {
	path := strings.Split(p.PathName, "/")

	if len(path) == 0 {
		return ""
	}
	rootFolder := path[0]
	return rootFolder
}

func (s *Store) Has(key string) bool {
	pathKey := s.pathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	_, err := os.Stat(fullPathWithRoot)
	return !errors.Is(err, os.ErrNotExist)
}

type StoreOpts struct {
	// Root is the base folder for all files and folders to be saved
	Root              string
	pathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		FileName: key,
	}
}

func NewStore(opts StoreOpts) *Store {
	if opts.pathTransformFunc == nil {
		opts.pathTransformFunc = DefaultPathTransformFunc
	}

	if len(opts.Root) == 0 {
		opts.Root = "home"
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Clear() error {
	return os.RemoveAll(s.Root)
}

func (s *Store) Delete(key string) error {
	pathKey := s.pathTransformFunc(key)

	defer func() {
		log.Printf("deleted [%s] from disk", pathKey.FileName)
	}()
	pathNameWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.RootFolder())
	return os.RemoveAll(pathNameWithRoot)
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	return buf, err
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.pathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	return os.Open(fullPathWithRoot)
}

func (s *Store) Write(key string, r io.Reader) error {
	return s.writeStream(key, r)
}

func (s *Store) writeStream(key string, r io.Reader) error {

	pathKey := s.pathTransformFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.PathName)
	if err := os.MkdirAll(pathNameWithRoot, os.ModePerm); err != nil {
		return err
	}

	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	f, err := os.Create(fullPathWithRoot)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("written %d bytes to disk: %s", n, pathKey.FileName)

	return nil
}
