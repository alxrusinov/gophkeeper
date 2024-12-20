package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (r *Router) GetCredential(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
