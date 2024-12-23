package httphandler

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) GetNote(ctx iris.Context) {
	fmt.Printf("USER - %#v", ctx.Values())
	ctx.StatusCode(http.StatusOK)
}
