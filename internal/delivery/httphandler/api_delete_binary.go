package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DeleteBinary - deletes binary data
func (h *HttpHandler) DeleteBinary(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
}
