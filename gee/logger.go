package gee

import (
	"log"
	"time"
)

// Logger 全局中间件
func Logger() HandlerFunc {
	return func(ctx *Context) {
		// Start timer
		t := time.Now()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group app", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func OnlyForV1() HandlerFunc {
	return func(ctx *Context) {
		// Start timer
		t := time.Now()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v1", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func OnlyForV2() HandlerFunc {
	return func(ctx *Context) {
		// Start timer
		t := time.Now()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
