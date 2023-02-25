package main

import (
	"gee"
	"net/http"
)

func main() {
	app := gee.New()

	app.Get("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	app.Get("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	app.Get("/header/:id", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "this match is %s, you're at %s\n", ctx.Param("id"), ctx.Path)
	})

	app.Post("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	app.Run(":8080")
}
