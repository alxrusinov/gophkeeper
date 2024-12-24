package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Logout - logiut handler
func (h *HttpHandler) Logout(ctx iris.Context) {
	ctx.RemoveCookie(userCookie)

	ctx.StatusCode(http.StatusUnauthorized)
	ctx.Writef("token invalidated, a new token is required to access the protected API")

}
