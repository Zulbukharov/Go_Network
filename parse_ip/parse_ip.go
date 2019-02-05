package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var ps *string

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <ip_address>\n", os.Args[0])
		os.Exit(0)
	}
	ps = &os.Args[1]
	addr := net.ParseIP(*ps)
	if addr == nil {
		fmt.Printf("Invalid address\n")
	} else {
		fmt.Printf("Address is %s\n", addr.String())
	}
	os.Exit(0)
}
