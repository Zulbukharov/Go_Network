/**
 * File              : lookup_port.go
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

//  func LookupPort(network, service string) (port int, err os.Error)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <network> <service>\n", os.Args[0])
		os.Exit(1)
	}
	port, err := net.LookupPort(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println("Lookup error: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Port using for [service][%s]: [port:][%d]\n", os.Args[2], port)
}
