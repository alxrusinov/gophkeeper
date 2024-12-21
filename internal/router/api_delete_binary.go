package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DeleteBinary - deletes binary data
func (r *Router) DeleteBinary(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
