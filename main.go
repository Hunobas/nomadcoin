package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	difficulty := 6
	target := strings.Repeat("0", difficulty)
	nonce := 1

	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
		if strings.HasPrefix(hash, target) {
			fmt.Printf("hello%d\n", nonce)
			fmt.Println(hash)
			return
		} else {
			nonce++
		}
	}
}
