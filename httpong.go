package httpong

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Addr  string
	Debug bool
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
	log.Println("running in port ", a.config.Addr)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
			}
			go handleConnection(conn, a.router)
		}
	}()

	<-sigCh

	log.Println("server shutting down..")
	return nil
}

func handleConnection(conn net.Conn, router Router) {
	defer conn.Close()

	req := ReadReq(conn)
	route := router.FindRoute(req)
	context := newContext(req, conn)
	if route == nil {
		context.res.ResponseString(conn, "failed to access path", 404)
		return
	}
	log.Println(route, context)
	route.Function(context)
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
