package httpong

import (
	"bytes"
	"fmt"
	"io"
)

type Res struct{}

func (r Res) ResponseString(conn io.ReadWriter, message string, statusCode int) error {
	output := new(bytes.Buffer)
	fmt.Fprintf(output, "HTTP/1.1 %d OK\r\n", statusCode)
	fmt.Fprintf(output, "Content-Type: text/plain\r\n")
	lenMessage := len(message)
	fmt.Fprintf(output, "Content-Length: %d\r\n", lenMessage)
	fmt.Fprintf(output, "\r\n")
	fmt.Fprintf(output, "%v", message)

	_, err := io.Copy(conn, output)
	return err
}
