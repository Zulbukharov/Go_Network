/**
 * File              : mask.go
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
		 fmt.Fprintf(os.Stderr, "Usage: %s <ip_address>\n", os.Args[0])
		 os.Exit(0)
	 }
	 dotAddr := os.Args[1]

	 addr := net.ParseIP(dotAddr)
	 if addr == nil {
		 fmt.Println("Invalid address")
		 os.Exit(0)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		" Leading ones count is ", ones,
		" Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}
