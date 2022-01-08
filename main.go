package main

import (
	"github.com/Hunobas/nomadcoin/blockchain"
	"github.com/Hunobas/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
