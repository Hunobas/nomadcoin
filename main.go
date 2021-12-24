package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("!!Welcome to 노마드 코인!!\n\n")
	fmt.Printf("\tPlease use the following commands:\n\n")
	fmt.Printf("\texplorer:	Start the HTML Explorer\n")
	fmt.Printf("\trest:		Start the REST API (recommanded)\n\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Set the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer...")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Printf("Start Server with the port %d\n", *portFlag)
	}
}
