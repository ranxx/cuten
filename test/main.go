package main

import (
	"cuten"
	"fmt"
	"net/http"
)

func main() {
	e := cuten.New()
	e.GET("/", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"axing": ctx.Path})
	})
	e.GET("/hello", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"阿星": ctx.Path})
	})
	e.GET("/axing", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"url": ctx.Path})
	})
	fmt.Println(e.Run(":9999"))
}
