package cuten

import (
	"net/http"
)

var join string = "-"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method, pattern string, handler HandlerFunc) {
	r.handlers[method+join+pattern] = handler
}

func (r *router) handle(ctx *Context) {
	handler, ok := r.handlers[ctx.Method+join+ctx.Path]
	if !ok {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
	handler(ctx)
}
