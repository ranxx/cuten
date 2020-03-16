package cuten

import (
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler use by cuten
type HandlerFunc func(ctx *Context)

// Engine implement the interface of http.Handler
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	// 找到这个路由所有组的handlers
	ctx.handlers = append(ctx.handlers, e.handlers...)
	for _, g := range e.groups {
		if strings.HasPrefix(ctx.Path, g.prefix) {
			ctx.handlers = append(ctx.handlers, g.handlers...)
		}
	}
	e.router.handle(ctx)
}

// New is constructor of cuten.Engine
func New() *Engine {
	e := &Engine{
		router: newRouter(),
		groups: make([]*RouterGroup, 0, 3),
	}
	e.RouterGroup = &RouterGroup{engine: e}
	return e
}

func (e *Engine) addRouter(method, pattern string, f HandlerFunc) {
	e.router.addRouter(method, pattern, f)
}

// // GET add get request
// func (e *Engine) GET(pattern string, f HandlerFunc) {
// 	e.addRouter("GET", pattern, f)
// }

// // POST add post request
// func (e *Engine) POST(pattern string, f HandlerFunc) {
// 	e.addRouter("POST", pattern, f)
// }

// Run start router
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
