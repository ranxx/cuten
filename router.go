package cuten

import (
	"fmt"
	"net/http"
	"strings"
)

var join string = "-"

type router struct {
	roots    map[string]*node
	handlers map[string][]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string][]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

// POST-/
// GET-/user/a
func (r *router) addRouter(method, pattern string, handler ...HandlerFunc) {
	if _, ok := r.handlers[method+join+pattern]; ok {
		// 已经注册过
		panic(fmt.Sprintf("%s already register !", method+join+pattern))
	}
	// 解析
	parts := strings.Split(pattern, "/")
	if len(parts) > 0 && parts[0] == "" {
		parts = parts[1:]
	}
	if len(parts) < 1 {
		// 路由有误
		panic(fmt.Sprintf("%s unknown path !", method+join+pattern))
	}
	// 第一层
	root, ok := r.roots[method+join+parts[0]]
	if !ok {
		root = &node{pattern: pattern, part: parts[0], precise: true}
		r.roots[method+join+parts[0]] = root
	}
	// 后面的层
	root.insert(pattern, parts[1:], 0)
	r.handlers[method+join+pattern] = handler
}

func notFound() HandlerFunc {
	return func(ctx *Context) {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
		ctx.Next()
	}
}

func (r *router) match(ctx *Context) []HandlerFunc {
	notFoundFs := func() []HandlerFunc {
		return []HandlerFunc{notFound()}
	}

	parts := strings.Split(ctx.Path, "/")
	if len(parts) > 0 && parts[0] == "" {
		parts = parts[1:]
	}
	if len(parts) <= 0 {
		return notFoundFs()
	}
	root, ok := r.roots[ctx.Method+join+parts[0]]
	if !ok {
		return notFoundFs()
	}

	n := root.search(parts[1:], 0)
	if n == nil {
		return notFoundFs()
	}
	handler, ok := r.handlers[ctx.Method+join+n.pattern]
	if !ok {
		return notFoundFs()
	}
	if strings.Index(n.pattern, ":") != -1 || strings.Index(n.pattern, "*") != -1 {
		ctx.parseURLParam(n.pattern)
	}
	return handler
}

func (r *router) handle(ctx *Context) {
	handler := r.match(ctx)
	ctx.handlers = append(ctx.handlers, handler...)
	ctx.Next()
}

type RouterGroup struct {
	prefix   string
	parent   *RouterGroup
	engine   *Engine
	handlers []HandlerFunc
}

func (g *RouterGroup) Use(f ...HandlerFunc) {
	g.handlers = append(g.handlers, f...)
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix:   g.prefix + prefix,
		parent:   g,
		engine:   engine,
		handlers: make([]HandlerFunc, 0, 3),
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRouter(method, pattern string, handler ...HandlerFunc) {
	pattern = g.prefix + pattern
	g.engine.addRouter(method, pattern, handler...)
}

func (g *RouterGroup) GET(pattern string, handler ...HandlerFunc) {
	g.addRouter("GET", pattern, handler...)
}

func (g *RouterGroup) POST(pattern string, handler ...HandlerFunc) {
	g.addRouter("POST", pattern, handler...)
}
