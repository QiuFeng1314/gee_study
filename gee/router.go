package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() (r *router) {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

// 解析pattern，将/p/:lang 分割成数组
func parsePattern(pattern string) (parts []string) {

	patterns := strings.Split(pattern, "/")

	for _, v := range patterns {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return
}

// 封装router数据
func (router *router) addRouter(method, pattern string, handler HandlerFunc) {

	parts := parsePattern(pattern)
	key := method + "-" + pattern

	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}

	router.roots[method].insert(pattern, parts, 0)
	router.handlers[key] = handler
}

func (router *router) getRouter(method string, path string) (n *node, params map[string]string) {

	searchParts := parsePattern(path)

	root, ok := router.roots[method]
	if ok {

		n = root.search(searchParts, 0)
		if n != nil {

			params = make(map[string]string)

			parts := parsePattern(n.pattern)
			for index, part := range parts {
				if part[0] == ':' {
					params[part[1:]] = searchParts[index]
				}
				if part[0] == '*' && len(part) > 1 {
					params[part[1:]] = strings.Join(searchParts[index:], "/")
					break
				}
			}
		}
	}

	return
}

func (router *router) handle(ctx *Context) {

	n, params := router.getRouter(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.pattern
		router.handlers[key](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
