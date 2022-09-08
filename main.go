package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(writer http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(10 << 20)     // memoria disponivel por requisição (10 mb)
	file, _, err := request.FormFile("file") // key da requisição
	if err != nil {
		fmt.Println("Erro ao receber o arquivo")
		fmt.Println(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("uploads", "file-*.png") // diretório e nome do arquivo
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(writer, "Upload feito com sucesso\n")
}

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
