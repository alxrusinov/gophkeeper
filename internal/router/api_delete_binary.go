package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (r *Router) DeleteBinary(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
