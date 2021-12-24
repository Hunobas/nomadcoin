package main

import (
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

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer...")
	case "rest":
		fmt.Println("Start REST API...")
	default:
		usage()
	}
}
