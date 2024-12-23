package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) GetBankCards(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
