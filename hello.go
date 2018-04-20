package main

import (
	"golang.org/x/crypto/nacl/box"
	"crypto/rand"
	"github.com/golang-demos/go-utils-demo/goutils"
	"fmt"
)

func main() {
	pubKey1, priKey1, _ := box.GenerateKey(rand.Reader)
	pubKey2, priKey2, _ := box.GenerateKey(rand.Reader)

	var sharedKey1 [32]byte
	box.Precompute(&sharedKey1, pubKey2, priKey1)

	var sharedKey2 [32]byte
	box.Precompute(&sharedKey2, pubKey1, priKey2)

	fmt.Println("shared key1 and key2 should be the same")
	goutils.PrintlnByteArray("shared key1", sharedKey1[:])
	goutils.PrintlnByteArray("shared key2", sharedKey2[:])
}
