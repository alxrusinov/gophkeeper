package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DeleteCredentials - deletes login and password data
func (r *Router) DeleteCredentials(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)

}
