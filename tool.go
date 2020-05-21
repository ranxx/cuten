package cuten

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// TerminalColor Terminal color
type TerminalColor string

var (
	nspac = []string{
		"",
		" ",
		"  ",
		"   ",
		"    ",
	}
)

// Terminal color
var (
	TCBlue    = TerminalColor("\033[1;36m")
	TCRed     = TerminalColor("\033[1;31m")
	TCGreen   = TerminalColor("\033[1;32m")
	TCYellow  = TerminalColor("\033[1;33m")
	TCDefault = TerminalColor("\033[0m")
)

func fileinfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return fmt.Sprintf("%s:%d", "<<<?>>>", 0)
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.String(http.StatusInternalServerError, "Internal Server Error\n")
			}
		}()
		ctx.Next()
	}
}

func reNSpace(method string) string {
	return nspac[7-len(method)]
}

func nowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05 ")
}

func debug(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func handlerSpendTime() HandlerFunc {
	start := time.Now()
	return func(ctx *Context) {
		defer func() {
			spendTime := time.Now().Sub(start)
			err := recover()
			if err != nil && ctx.StatusCode == http.StatusOK {
				ctx.StatusCode = http.StatusInternalServerError
			}
			if ctx.StatusCode == http.StatusOK {
				debug("%s%s [%d] %s%s %s in %v%s\n", TCGreen, nowTimeString(), ctx.StatusCode, ctx.Method, reNSpace(ctx.Method), ctx.Path, spendTime, TCDefault)
			} else if err != nil {
				debug("%s%s [%d] %s%s %s %s in %v%s\n", TCRed, nowTimeString(), ctx.StatusCode, ctx.Method, reNSpace(ctx.Method), ctx.Path, err, spendTime, TCDefault)
			} else if ctx.StatusCode != http.StatusOK {
				debug("%s%s [%d] %s%s %s in %v%s\n", TCYellow, nowTimeString(), ctx.StatusCode, ctx.Method, reNSpace(ctx.Method), ctx.Path, spendTime, TCDefault)
			}
			if err != nil {
				panic(err)
			}
		}()
		ctx.Next()
	}
}
