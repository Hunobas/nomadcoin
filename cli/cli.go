package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Hunobas/nomadcoin/explorer"
	"github.com/Hunobas/nomadcoin/rest"
)

func usage() {
	fmt.Printf("!!Welcome to 노마드 코인!!\n\n")
	fmt.Printf("\tPlease use the following flags:\n\n")
	fmt.Printf("\t-port:	Set PORT of the server\n")
	fmt.Printf("\t-mode:	Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	port := flag.Int("port", 4000, "Set port of the server")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(port)
	fmt.Printf("Start Server with the port %d in %s\n", *port, *mode)
}
