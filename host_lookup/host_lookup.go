/**
 * File              : host_lookup.go
 * Author            : azulbukh <l1pt0n1905@gmail.com>
 * Date              : 28.01.2019
 * Last Modified Date: 28.01.2019
 * Last Modified By  : azulbukh <l1pt0n1905@gmail.com>
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <name>", os.Args[0])
		os.Exit(1)
	}
	addrs, err := net.LookupHost(os.Args[1])
	if err != nil {
		fmt.Println("Lookup error:", err.Error())
		os.Exit(1)
	}
	fmt.Println(os.Args[1], " host addresses:")
	for _, ip := range addrs {
		fmt.Println("	", ip)
	}
	os.Exit(0)
}
