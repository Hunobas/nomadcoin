package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

func FromeBytes(i interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}

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
