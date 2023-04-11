package main

import (
	"gee/engine"
	"net/http"
)

func main() {
	r := engine.New()
	r.GET("/", func(ctx *engine.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(ctx *engine.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	r.POST("/login", func(ctx *engine.Context) {
		ctx.JSON(http.StatusOK, engine.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.Run(":9999")
}
