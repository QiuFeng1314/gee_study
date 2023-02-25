package gee

import (
	"fmt"
	"net/http"
	"util"
)

type H map[string]any

type HandlerFunc func(*Context)

type Context struct {
	Resp       http.ResponseWriter
	Req        *http.Request
	Path       string // 请求路径
	Method     string // 请求类型
	StatusCode int    // 状态码
}

func newContext(resp http.ResponseWriter, req *http.Request) (ctx *Context) {
	return &Context{
		Resp:   resp,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Resp.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Resp.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, values ...any) {
	ctx.SetHeader("Context-Type", "text/plain")
	ctx.Status(code)
	ctx.Resp.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, obj any) {
	ctx.SetHeader("Context-Type", "application/json")
	ctx.Status(code)
	encoder := util.Json.NewEncoder(ctx.Resp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Resp, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Resp.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Context-Type", "text/html")
	ctx.Status(code)
	ctx.Resp.Write([]byte(html))
}
