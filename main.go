package main

import (
	"github.com/Hunobas/nomadcoin/cli"
	"github.com/Hunobas/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
