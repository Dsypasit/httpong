# httpong: simple http library

I made it for learning how http framework work. I know my codebase was a bit weird but i think i learn
a lot from this project. I learned about TDD and how http work at the TCP layer (a little bit :))

## Example

```golang

package main

import (
	"fmt"

	"github.com/Dsypasit/httpong"
)

func main() {
	config := httpong.Config{
		Addr: ":8080",
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

```

## Goal

- [ ] implement cookie
- [ ] make middleware
- more...

## Inspiration

- [net](https://pkg.go.dev/net)
- [echo](https://github.com/labstack/echo)
