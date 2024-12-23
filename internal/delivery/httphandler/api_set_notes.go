package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetNote(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
