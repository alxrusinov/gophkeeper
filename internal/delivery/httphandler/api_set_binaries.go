package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetBinry(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
