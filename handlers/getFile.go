package handlers

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetFile(context *routing.Context) error {
	fmt.Fprintf(context, "Hello world")
	return nil
}
