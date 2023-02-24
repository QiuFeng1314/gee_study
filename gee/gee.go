package gee

import (
	"log"
	"net/http"
)

var (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() (engine *Engine) {
	return &Engine{router: newRouter()}
}

func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	engine.router.handle(newContext(resp, req))
}

func (engine *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRouter(GET, pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRouter(POST, pattern, handler)
}

func (engine *Engine) Put(pattern string, handler HandlerFunc) {
	engine.addRouter(PUT, pattern, handler)
}

func (engine *Engine) Delete(pattern string, handler HandlerFunc) {
	engine.addRouter(DELETE, pattern, handler)
}

func (engine *Engine) Run(port string) (err error) {
	err = http.ListenAndServe(port, engine)
	if err == nil {
		log.Printf("gee start success! server port is %q...", port)
	}
	return
}
