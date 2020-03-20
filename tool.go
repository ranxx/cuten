package cuten

import (
	"fmt"
	"net/http"
	"runtime"
)

func fileinfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return fmt.Sprintf("%s:%d", "<<<?>>>", 0)
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func Recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.String(http.StatusInternalServerError, "Internal Server Error\n")
			}
		}()
		ctx.Next()
	}
}
