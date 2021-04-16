package plugins

import (
	"context"
	"net/http"
)

type Context struct {
	context.Context
	Request *http.Request
	Writer http.ResponseWriter

	index    int
	handlers []HandlerFunc
}

type HandlerFunc func(c *Context)

func (c *Context) Next() {
	c.index++
	c.handlers[c.index](c)
}
