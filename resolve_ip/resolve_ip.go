/**
 * File              : resolve_ip.go
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
	var namePointer *string

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <name>\n", os.Args[0])
		os.Exit(1)
	}
	namePointer = &os.Args[1]
	// resolve ip addr accept ip or ip4 or ip6 and name of website
	addr, err := net.ResolveIPAddr("ip", *namePointer)
	if err != nil {
		fmt.Println("Resolution error ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolution address ", addr.String())
	os.Exit(0)
}
