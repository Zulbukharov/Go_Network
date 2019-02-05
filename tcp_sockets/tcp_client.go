package main

/**
 * File              : tcp_client.go
 * Author            : azulbukh <l1pt0n1905@gmail.com>
 * Date              : 04.02.2019
 * Last Modified Date: 04.02.2019
 * Last Modified By  : azulbukh <l1pt0n1905@gmail.com>
 */

// func (c *TCPConn) Write(b []byte) (n int, err os.Error)
// func (c *TCPConn) Read(b []byte) (n int, err os.Error)
// type TCPAddr struct {
//     IP   IP
//     Port int
// }
// func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)

// creates connection between two sockets to send and recieve data
// func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
// laddr == nil
// raddr remote TCPAddr
// net tcp, tcp4, tcp6

// go run tcp_client.go www.google.com:80
// [ HTTP/1.0 200 OK
// Date: Mon, 04 Feb 2019 11:39:08 GMT
// Expires: -1
// Cache-Control: private, max-age=0
// Content-Type: text/html; charset=ISO-8859-1
// P3P: CP="This is not a P3P policy! See g.co/p3phelp for more info."
// Server: gws
// X-XSS-Protection: 1; mode=block
// X-Frame-Options: SAMEORIGIN
// Set-Cookie: 1P_JAR=2019-02-04-11; expires=Wed, 06-Mar-2019 11:39:08 GMT; path=/; domain=.google.com
// Set-Cookie: NID=158=VIlO7QIuMI_8V-olYnJtmlavFDV-NSDU98iza7DC7z9N64ZOS1QWabX-43Qqr8EYpB3deoNp66UI4T_of4UuVN8XiFQVO4pUdtLAV891GJR9QjGlW1_94rBPMcc1hCbXjZ5-vc4QihUOnD1rgtOYQHUDIe7HzlqWfrTjI6bjk0c; expires=Tue, 06-Aug-2019 11:39:08 GMT; path=/; domain=.google.com; HttpOnly
// Accept-Ranges: none
// Vary: Accept-Encoding
//  ]

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <remote address>:<port>", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println("[", string(result), "]")
	os.Exit(0)
}
