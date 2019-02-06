/**
 * File              : ftp_server.go
 * Author            : azulbukh <l1pt0n1905@gmail.com>
 * Date              : 06.02.2019
 * Last Modified Date: 06.02.2019
 * Last Modified By  : azulbukh <l1pt0n1905@gmail.com>
 */

// This example deals with a directory browsing protocol -
// basically a stripped down version of FTP, but without
// even the file transfer part. We only consider listing
// a directory name, listing the contents of a directory
// and changing the current directory - all on the server
// side, of course. This is a complete worked example of
// creating all components of a client-server application.
// It is a simple program which includes messages in both
// directions, as well as design of messaging protocol.

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func chdir(conn net.Conn, s string) {
	defer conn.Write([]byte("\r\n"))
	if os.Chdir(s[0:len(s)-2]) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))

}

func dirList(conn net.Conn) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return
	}
	for _, f := range files {
		conn.Write([]byte("\033[1;32" + f.Name() + "\r\n\033[0m"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FATAL ERROR][%s]\n", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}
		s := string(buf[0:n])
		if s[0:2] == CD {
			chdir(conn, s[3:])
		} else if s[0:3] == DIR {
			dirList(conn)
		} else if s[0:3] == PWD {
			pwd(conn)
		}
	}
}

func main() {
	service := "0.0.0.0:1202"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
