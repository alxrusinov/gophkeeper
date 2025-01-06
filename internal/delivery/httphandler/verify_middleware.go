package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) VerifyMiddleware(ctx iris.Context) {
	userFromCtx, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	checkedUser, err := h.usecase.CheckUser(ctx, userFromCtx.ID)

	if err != nil || !checkedUser {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()

}
