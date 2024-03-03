package main

import (
	"net"

	"github.com/Dsypasit/httpong"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	httpong.ReadReq(conn)
	httpong.ResponseString(conn, "hello world", 200)
}
