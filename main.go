package main

import (
	"github.com/matheusgb/pocPostFile/handlers"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	router := routing.New()

	router.Get("/upload", handlers.GetFile)

	fasthttp.ListenAndServe(":8080", router.HandleRequest)
}
