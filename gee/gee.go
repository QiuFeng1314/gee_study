package gee

import (
	"fmt"
	"log"
	"net/http"
)

var (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type HandlerFunc func(resp http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(resp, req)
	} else {
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL)
	}
}

func New() (engine *Engine) {
	return &Engine{router: make(map[string]HandlerFunc, 4)}
}

func (engine *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
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
