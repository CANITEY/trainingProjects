package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/md4"
	"hash"
	"io"
)

func hasher(str string, hash hash.Hash) string {
	hasher := hash
	io.WriteString(hasher, str)
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	fmt.Printf("md5 hash: %v\n", hasher("hi there", md5.New()))
	fmt.Printf("md4 hash: %v\n", hasher("hi there", md4.New()))
	fmt.Printf("sha1 hash: %v\n", hasher("hi there", sha1.New()))
	fmt.Printf("sha256 hash: %v\n", hasher("hithere", sha256.New()))
}
