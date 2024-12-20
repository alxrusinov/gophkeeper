package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (r *Router) GetBankCards(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
