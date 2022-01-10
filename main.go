package main

import (
	"github.com/Hunobas/nomadcoin/blockchain"
	"github.com/Hunobas/nomadcoin/cli"
	"github.com/Hunobas/nomadcoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
