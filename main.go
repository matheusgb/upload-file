package main

import (
	"fmt"

	routing "github.com/rmnoff/fasthttp-routing/v3"
	"github.com/valyala/fasthttp"
)

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

func main() {
	router := routing.New()

	router.Post("/upload", uploadFile)
	fasthttp.ListenAndServe(":8080", router.HandleRequest)
}
