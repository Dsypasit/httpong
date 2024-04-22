package httpong

import (
	"fmt"
	"io"
	"net"
	"regexp"
	"strconv"
)

type Req struct {
	Body        string
	ContentType string
	Method      string
}

// read request from client
func ReadReq(conn net.Conn) (req Req) {
	// var arr []byte
	size := 4 << 10
	a := make([]byte, size)
	// conn.SetDeadline(time.Now().Add(1 * time.Second))
	// conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	_, err := conn.Read(a)
	if err == io.EOF {
		fmt.Println("break")
	}
	// fmt.Println("read", readBytes)
	// fmt.Printf("%v\n", string(a))
	fmt.Println("method", getHTTPMethod(string(a)))
	fmt.Println("content-length", getContentLength(string(a)))
	fmt.Println("body", getBody(string(a)))

	req = Req{
		Body:        getBody(string(a)),
		Method:      getHTTPMethod(string(a)),
		ContentType: getContentType(string(a)),
	}
	return req
}

// return http method
func getHTTPMethod(l string) string {
	pattern := `^(\w+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindString(l)
	return matches
}

func getContentLength(l string) int {
	pattern := `Content-Length: (\d+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(l)
	if len(matches) == 0 {
		return 0
	}
	length, _ := strconv.Atoi(matches[1])
	return length
}

func getContentType(l string) string {
	pattern := `Content-Type: (\w+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(l)
	if len(matches) == 0 {
		return ""
	}
	return matches[1]
}

func getBody(l string) string {
	pattern := `\r\n\r\n(.*)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(l)
	if len(matches) == 0 {
		return ""
	}
	return matches[1]
}
