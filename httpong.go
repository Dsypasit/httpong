package httpong

import (
	"fmt"
	"net"
)

type Config struct {
	Addr string
}

type App struct {
	router Router
	config Config
}

func New() App {
	return App{
		router: newRouter(),
		config: Config{Addr: ":8080"},
	}
}

func NewWithConfig(config Config) App {
	return App{
		router: newRouter(),
		config: config,
	}
}

func (a *App) Run() error {
	ln, err := net.Listen("tcp", a.config.Addr)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	ReadReq(conn)
	ResponseString(conn, "hello world", 200)
}

func (a *App) RegisterRoute(method string, path string, fn Handler) {
	route := Route{method, path, fn}
	a.router.registerRoute(route)
}

func (a *App) GET(path string, fn Handler) {
	a.RegisterRoute("GET", path, fn)
}

func (a App) POST(path string, fn Handler) {
	a.RegisterRoute("POST", path, fn)
}

func (a App) PUT(path string, fn Handler) {
	a.RegisterRoute("PUT", path, fn)
}

func (a App) DELETE(path string, fn Handler) {
	a.RegisterRoute("DELETE", path, fn)
}

func (a App) PATCH(path string, fn Handler) {
	a.RegisterRoute("PATCH", path, fn)
}
