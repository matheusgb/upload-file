package main

import (
	"fmt"
	"math"

	routing "github.com/rmnoff/fasthttp-routing/v3"
	"github.com/valyala/fasthttp"
)

func main() {
	router := routing.New()

	server := &fasthttp.Server{
		Handler:            router.HandleRequest,
		Name:               "FastHTTP Server",
		MaxRequestBodySize: math.MaxInt32,
	}

	router.Post("/upload", uploadFile)
	server.ListenAndServe(":8080")

}

func uploadFile(ctx *routing.Context) error {
	fileHeader, err := ctx.FormFile("file") // key da requisição
	if err != nil {
		fmt.Println("Erro ao receber o arquivo")
		return nil
	}

	fasthttp.SaveMultipartFile(fileHeader, fmt.Sprintf("./uploads/%s", fileHeader.Filename))

	fmt.Println("Arquivo recebido com sucesso")
	return nil
}
