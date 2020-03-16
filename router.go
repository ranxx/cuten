package cuten

import (
	"net/http"
	"strings"
)

var join string = "-"

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

// POST-/
// GET-/user/a
func (r *router) addRouter(method, pattern string, handler HandlerFunc) {
	if _, ok := r.handlers[method+join+pattern]; ok {
		// 已经注册过
		return
	}
	if pattern != "/" {
		// 解析
		parts := strings.Split(pattern, "/")
		if len(parts) > 0 && parts[0] == "" {
			parts = parts[1:]
		}
		if len(parts) < 1 {
			// 路由有误
			return
		}
		// 第一层
		root, ok := r.roots[method+join+parts[0]]
		if !ok {
			root = &node{pattern: pattern, part: parts[0], precise: true}
			r.roots[method+join+parts[0]] = root
		}
		root.insert(pattern, parts[1:], 0)
	}
	r.handlers[method+join+pattern] = handler
}

func (r *router) handle(ctx *Context) {
	// fmt.Printf("handlers:\t%#v\n", r.handlers)
	// fmt.Printf("roots:\t%#v\n", r.roots)
	register := ctx.Path
	if ctx.Path != "/" {
		parts := strings.Split(ctx.Path, "/")
		if len(parts) > 0 && parts[0] == "" {
			parts = parts[1:]
		}
		if len(parts) <= 0 {
			// 没有路由
			ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
			return
		}
		// 第一层
		root, ok := r.roots[ctx.Method+join+parts[0]]
		if !ok {
			ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
			return
		}
		// 匹配到的注册的路由信息
		n := root.search(parts[1:], 0)
		if n == nil {
			ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
			return
		}
		register = n.pattern
	}
	handler, ok := r.handlers[ctx.Method+join+register]
	if !ok {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
		return
	}

	// 处理参数
	if strings.Index(register, ":") != -1 {
		ctx.parseURLParam(register)
	}

	handler(ctx)
}

type RouterGroup struct {
	prefix string
	parent *RouterGroup
	engine *Engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRouter(method, pattern string, handler HandlerFunc) {
	pattern = g.prefix + pattern
	g.engine.router.addRouter(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRouter("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRouter("POST", pattern, handler)
}
