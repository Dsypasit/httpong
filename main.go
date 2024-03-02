package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
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

	responseString(conn, "hello world", 200)
}

func responseString(conn io.ReadWriter, message string, statusCode int) {
	output := new(bytes.Buffer)
	fmt.Fprintf(output, "HTTP/1.1 %d OK\r\n", statusCode)
	fmt.Fprintf(output, "Content-Type: text/plain\r\n")
	lenMessage := len(message)
	fmt.Fprintf(output, "Content-Length: %d\r\n", lenMessage)
	fmt.Fprintf(output, "\r\n")
	fmt.Fprintf(output, "%v", message)

	io.Copy(conn, output)
}
