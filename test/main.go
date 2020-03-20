package main

import (
	"cuten"
	"fmt"
	"net/http"
	"time"
)

var format string = "2006-01-02 15:04:05"

func timeWithoutHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

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
		fmt.Printf("我是%s路由中间件\n", ctx.Path)
	})
	v1.Use(func(ctx *cuten.Context) {
		fmt.Printf("我是%s路由中间件\n", ctx.Path)
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
	e.POST("/login/:user/info", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "user": ctx.URLParam["user"], "passwd": ctx.PostForm("passwd"), "info": "nil"})
	})
	e.GET("/login", func(ctx *cuten.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "user": ctx.Quary("user"), "passwd": ctx.Quary("passwd")})
	})
	file := e.Group("/file")
	{
		file.GET("/file/*filename", func(ctx *cuten.Context) {
			ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "file-name": ctx.URLParam["filename"]})
		})
		// // panic
		// e.GET("/file/*filename", func(ctx *cuten.Context) {
		// 	ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "filename": ctx.URLParam["filename"]})
		// })
		// // panic
		// e.GET("/file/css/*filename", func(ctx *cuten.Context) {
		// 	ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "css-filename": ctx.URLParam["filename"]})
		// })
		file.GET("/css/*filename", func(ctx *cuten.Context) {
			ctx.JSON(http.StatusOK, map[string]string{"URL": ctx.Path, "css-name": ctx.URLParam["filename"][100:]})
		})
	}

	fmt.Println(e.Run(":9999"))
}
