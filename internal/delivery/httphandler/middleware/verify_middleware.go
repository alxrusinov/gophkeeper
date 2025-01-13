package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (m *Middleware) VerifyMiddleware(ctx iris.Context) {
	userFromCtx, err := m.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	checkedUser, err := m.usecase.CheckUser(ctx, userFromCtx.ID)

	if err != nil || !checkedUser {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()

}
