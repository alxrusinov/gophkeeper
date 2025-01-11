package httphandler

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func (h *HttpHandler) BodyLimitMiddleware(limitBytes int64) context.Handler {
	return iris.LimitRequestBodySize(limitBytes)

}
