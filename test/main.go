package main

import (
	"cuten"
	"fmt"
	"net/http"
)

func main() {
	e := cuten.New()
	v1 := e.Group("/v1")
	{
		v1.GET("", func(ctx *cuten.Context) {
			ctx.String(http.StatusOK, "/v1\n")
		})
		v1.GET("/user", func(ctx *cuten.Context) {
			ctx.String(http.StatusOK, "%s\n", ctx.Path)
		})
	}
	e.Use(func(ctx *cuten.Context) {
		fmt.Printf("我是/路由中间件\n")
	})
	v1.Use(func(ctx *cuten.Context) {
		fmt.Printf("我是/v1路由中间件\n")
	})
	e.GET("/", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"axing": ctx.Path})
	})
	e.GET("/hello", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"阿星": ctx.Path})
	})
	e.GET("/axing", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"url": ctx.Path})
	})
	e.POST("/login", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "user": ctx.PostForm("user"), "passwd": ctx.PostForm("passwd")})
	})
	e.POST("/login/:user", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "user": ctx.URLParam["user"], "passwd": ctx.PostForm("passwd")})
	})
	e.GET("/login", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "user": ctx.Quary("user"), "passwd": ctx.Quary("passwd")})
	})

	fmt.Println(e.Run(":9999"))
}
