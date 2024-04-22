package httpong

import "net"

type Context struct {
	req  Req
	res  Res
	conn net.Conn
}

func newContext(req Req, conn net.Conn) *Context {
	return &Context{req, Res{}, conn}
}

func (c *Context) Send(statusCode int, msg string) error {
	return c.res.ResponseString(c.conn, msg, statusCode)
}
