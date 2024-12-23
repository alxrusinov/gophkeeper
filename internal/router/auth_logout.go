package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Logout - logiut handler
func (r *Router) Logout(ctx iris.Context) {
	err := ctx.Logout()

	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(err.Error())
	} else {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.Writef("token invalidated, a new token is required to access the protected API")
	}

	ctx.StatusCode(http.StatusOK)
}
