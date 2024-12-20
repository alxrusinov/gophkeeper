package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Register- method of registration new user
func (r *Router) Register(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
