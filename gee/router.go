package gee

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() (r *router) {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (router *router) addRouter(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	router.handlers[key] = handler
}

func (router *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := router.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
