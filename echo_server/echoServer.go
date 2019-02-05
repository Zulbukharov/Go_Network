/**
 * File              : echoServer.go
 * Author            : azulbukh <l1pt0n1905@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : azulbukh <l1pt0n1905@gmail.com>
 */

//"echo" is simple IETF service. This just reads what the client types, and sends it back:
package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR][%s]", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
