package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func (m *Middleware) BodyLimitMiddleware(limitBytes int64) context.Handler {
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
