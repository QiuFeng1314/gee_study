package main

import (
	"gee"
	"net/http"
)

func main() {
	app := gee.Default()
	{
		// 加载文件路径
		app.LoadHTMLGlob("templates/*")
		// 自定义函数，提供给前端使用
		//app.SetFuncMap(template.FuncMap{
		//	"Date": func() {
		//
		//	},
		//})
		app.Get("/", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "custom.tmpl", gee.H{
				"title": "gee",
				"name":  "景帅",
			})
		})

		app.Get("/panic", func(ctx *gee.Context) {
			names := []string{"ok"}
			ctx.JSON(http.StatusOK, names[2])
		})

		app.Static("/assets", "./static")
	}

	v1 := app.Group("/v1")
	v1.Use(gee.OnlyForV1())
	{
		v1.Get("/", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})

		v1.Get("/header/:id", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "this match is %s, you're at %s\n", ctx.Param("id"), ctx.Path)
		})
	}

	v2 := app.Group("/v2")
	v2.Use(gee.OnlyForV2())
	{
		v2.Get("/", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", "v2", ctx.Path)
		})

		v2.Post("/login", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})

		v2.Get("/assets/*filepath", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{"filepath": ctx.Param("filepath")})
		})
	}

	app.Run(":8080")
}
