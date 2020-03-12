package cuten

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type H map[string]interface{}

type Context struct {
	Writer      http.ResponseWriter
	Req         *http.Request
	Path        string
	Method      string
	StatuscCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:      w,
		Req:         r,
		Path:        r.URL.Path,
		Method:      r.Method,
		StatuscCode: 200,
	}
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Quary(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatuscCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	if err := json.NewEncoder(ctx.Writer).Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), code)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}
