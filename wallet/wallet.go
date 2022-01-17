package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/Hunobas/nomadcoin/utils"
)

const (
	hashedMsg  string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey string = "30770201010420a49664076a11c8952dfca0d53734e2c7b26d5278edc75eaa3f3a3916bd3be0e7a00a06082a8648ce3d030107a144034200045b392561680999303542c37988c6a2e93ec44cd519f776fcd5fbac5075a581608da5514ed17de8c1c4d22c3b4085bac2ef2ef8ec884222ac4f62dbde929eab39"
	signature  string = "7a4e72a4116755a59a5950cbe70fbc8e8a7f8fc39ea6bb358f6380770286aed71ae46e260ef456ae775c667ee18d326a588d59705bfdb608f9044ddb85ca18ca"
)

func Start() {
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)
	_, err = x509.ParseECPrivateKey(privByte)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]
	fmt.Printf("%d\n\n%d\n\n%d\n\n", sigBytes, rBytes, sBytes)

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
