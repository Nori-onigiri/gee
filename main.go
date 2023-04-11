package main

import (
	"gee/engine"
	"net/http"
)

func main() {
	r := engine.New()
	r.GET("/", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *engine.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *engine.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *engine.Context) {
		c.JSON(http.StatusOK, engine.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
