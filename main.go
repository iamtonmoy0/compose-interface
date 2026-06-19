package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {

	data := []byte("hello world")
	hashAndBroadcast(bytes.NewReader(data))
}

func hashAndBroadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	hash := sha1.Sum(b)
	fmt.Println(hex.EncodeToString(hash[:]))
	return broadcast(r)
}
func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the byets:", string(b))
	return nil
}
