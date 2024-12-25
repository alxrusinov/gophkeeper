package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetBinary(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
