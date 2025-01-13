package authhandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Logout - logout handler
func (ah *AuthHandler) Logout(ctx iris.Context) {
	ctx.RemoveCookie(ah.userCookie)

	ctx.StatusCode(http.StatusUnauthorized)
	ctx.Writef("token invalidated, a new token is required to access the protected API")

}
