package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetBankCards(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
