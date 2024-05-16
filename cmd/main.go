package main

import (
	"fmt"

	"github.com/Dsypasit/httpong"
)

func main() {
	config := httpong.Config{
		Addr: ":8000",
	}
	app := httpong.NewWithConfig(config)

	app.GET("/", func(ctx *httpong.Context) error {
		return ctx.Send(200, "hello world")
	})

	err := app.Run()
	if err != nil {
		fmt.Printf("failed to run: %v", err)
	}
}
