package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("!!Welcome to 노마드 코인!!\n\n")
	fmt.Printf("\tPlease use the following commands:\n\n")
	fmt.Printf("\t-mode=html:\tStart the HTML Explorer\n")
	fmt.Printf("\t-mode=rest:\tStart the REST API (recommanded)\n\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	flags := flag.NewFlagSet("rest", flag.ExitOnError)
	modeFlag := flags.String("mode", "html", "Set the mode of the server")
	portFlag := flags.Int("port", 4000, "Set the port of the server")

	if flags.Parsed() {
		fmt.Println(portFlag)
		fmt.Printf("Start Server with the port %d\n in %s", *portFlag, *modeFlag)
	}
}
