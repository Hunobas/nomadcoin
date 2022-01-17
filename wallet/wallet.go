package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/Hunobas/nomadcoin/utils"
)

const (
	hashedMsg  string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey string = "30770201010420a49664076a11c8952dfca0d53734e2c7b26d5278edc75eaa3f3a3916bd3be0e7a00a06082a8648ce3d030107a144034200045b392561680999303542c37988c6a2e93ec44cd519f776fcd5fbac5075a581608da5514ed17de8c1c4d22c3b4085bac2ef2ef8ec884222ac4f62dbde929eab39"
	signature  string = "7a4e72a4116755a59a5950cbe70fbc8e8a7f8fc39ea6bb358f6380770286aed71ae46e260ef456ae775c667ee18d326a588d59705bfdb608f9044ddb85ca18ca"
)

func Start() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	utils.HandleErr(err)
	fmt.Printf("%x\n", keyAsBytes)

	hashAsBytes, err := hex.DecodeString(hashedMsg)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	signature := append(r.Bytes(), s.Bytes()...)

	fmt.Printf("%x\n", signature)
	utils.HandleErr(err)

}
