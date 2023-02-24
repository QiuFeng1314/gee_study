package main

import (
	"gee"
	"net/http"
)

func main() {
	app := gee.New()

	app.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	app.Get("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	app.Run(":8080")
}
