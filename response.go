package httpong

import (
	"bytes"
	"fmt"
	"io"
)

type Res struct {
	Status  int
	Message string
}

func ResponseString(conn io.ReadWriter, message string, statusCode int) {
	output := new(bytes.Buffer)
	fmt.Fprintf(output, "HTTP/1.1 %d OK\r\n", statusCode)
	fmt.Fprintf(output, "Content-Type: text/plain\r\n")
	lenMessage := len(message)
	fmt.Fprintf(output, "Content-Length: %d\r\n", lenMessage)
	fmt.Fprintf(output, "\r\n")
	fmt.Fprintf(output, "%v", message)

	io.Copy(conn, output)
}
