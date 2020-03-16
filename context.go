package cuten

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// type H map[string]interface{}

type Context struct {
	Writer      http.ResponseWriter
	Req         *http.Request
	Path        string
	Method      string
	StatuscCode int
	URLParam    map[string]string
	handlers    []HandlerFunc
	index       int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:      w,
		Req:         r,
		Path:        r.URL.Path,
		Method:      r.Method,
		StatuscCode: 200,
		URLParam:    make(map[string]string),
		handlers:    make([]HandlerFunc, 0, 3),
		index:       -1,
	}
}

// PostForm return the value of the HTTPbody
func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

// Quary return value thee URL
func (ctx *Context) Quary(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Param(key string) {

}

// Status HTTP status code
func (ctx *Context) Status(code int) {
	ctx.StatuscCode = code
	ctx.Writer.WriteHeader(code)
}

// SetHeader resp hander
func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

// String resp message type
func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON resp message type
func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	if err := json.NewEncoder(ctx.Writer).Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), code)
	}
}

// Data resp message type
func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Writer.Write(data)
}

// HTML resp message type
func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}

func (ctx *Context) parseURLParam(register string) {
	p1 := strings.Split(register, "/")
	p2 := strings.Split(ctx.Path, "/")
	for i, v := range p1 {
		if len(v) > 0 && v[0] == ':' {
			ctx.URLParam[v[1:]] = p2[i]
		}
	}
}

func (ctx *Context) Next() {
	ctx.index++
	s := len(ctx.handlers)
	for ; ctx.index < s; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}
