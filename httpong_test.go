package httpong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	app := New()
	assert.Equal(t, ":8080", app.config.Addr)
}

func TestNewWithConfig(t *testing.T) {
	config := Config{Addr: ":8081"}
	app := NewWithConfig(config)

	assert.Equal(t, ":8081", app.config.Addr)
}

func TestGet(t *testing.T) {
	app := New()

	app.GET("/hello", func(ctx *Context) error {
		return nil
	})

	assert.Equal(t, 1, len(app.router.route))
	assert.Equal(t, "/hello", app.router.route[0].Path)
}
