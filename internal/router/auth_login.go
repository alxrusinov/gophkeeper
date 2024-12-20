package router

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Login - method of loginng existing user
func (r *Router) Login(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
