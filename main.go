package main

import (
	"gee/engine"
	"net/http"
)

func main() {
	r := engine.Default()
	r.GET("/", func(c *engine.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *engine.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
