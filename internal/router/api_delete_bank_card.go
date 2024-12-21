package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DeleteBankCard - deletes bank card
func (r *Router) DeleteBankCard(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
