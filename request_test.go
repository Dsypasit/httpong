package httpong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReqPath(t *testing.T) {
	str := `POST / HTTP/1.1
Host: localhost:8080
User-Agent: curl/8.4.0
Accept: */*
Content-Length: 167
Content-Type: application/x-www-form-urlencoded

hello`
	expected := "/"
	actual := getPath(str)
	assert.Equal(t, expected, actual)
}
