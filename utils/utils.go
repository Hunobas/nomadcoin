// Package utils contains functions to be used across the application.
package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var logFn = log.Panic

func HandleErr(err error) {
	if err != nil {
		logFn(err)
	}
}

func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

// Frombytes takes an interface and data, then will encode the data to the  interface
func FromeBytes(i interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}

// Hash takes an interface, hashes it and returns the hex encoding of the hash.
func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func DecodeStringOrErr(s string) (bytes []byte) {
	bytes, err := hex.DecodeString(s)
	HandleErr(err)
	return
}

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}

func ToJSON(i interface{}) []byte {
	r, err := json.Marshal(i)
	HandleErr(err)
	return r
}
