package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func (h *HttpHandler) BodyLimitMiddleware(limitBytes int64) context.Handler {
	return func(ctx iris.Context) {
		method := ctx.Method()
		length := ctx.Request().ContentLength

		ctx.SetMaxRequestBodySize(limitBytes)

		if method == http.MethodPost && length > limitBytes {
			ctx.StopWithStatus(http.StatusRequestEntityTooLarge)
			return
		}

		ctx.Next()
	}

}
