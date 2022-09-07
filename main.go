package main

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	router := routing.New()

	router.Get("/upload", func(context *routing.Context) error {
		fmt.Fprintf(context, "Hello, world!")
		return nil
	})

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}
