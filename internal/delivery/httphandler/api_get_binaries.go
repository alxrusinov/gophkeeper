package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) GetBinaries(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}