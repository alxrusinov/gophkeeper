package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DeleteBankCard - deletes bank card
func (h *HttpHandler) DeleteBankCard(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
