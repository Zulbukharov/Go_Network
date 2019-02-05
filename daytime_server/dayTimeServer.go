/**
 * File              : dayTimeServer.go
 * Author            : azulbukh <l1pt0n1905@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : azulbukh <l1pt0n1905@gmail.com>
 */

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
