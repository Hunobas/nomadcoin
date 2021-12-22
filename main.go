package main

import (
	"github.com/Hunobas/nomadcoin/explorer"
	"github.com/Hunobas/nomadcoin/rest"
)

func main() {
	go explorer.Start(4000)
	rest.Start(4000)
}
