package main

import (
	"fmt"
	"gee"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	app := gee.New()

	app.Get("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "hello, this path is %v", req.URL.Path)
	})

	app.Get("/header", func(resp http.ResponseWriter, req *http.Request) {
		maps := make(map[string]any)
		for k, v := range req.Header {
			maps[k] = v
		}
		b, _ := json.Marshal(maps)
		fmt.Fprintf(resp, string(b))
	})

	app.Run(":8080")
}
