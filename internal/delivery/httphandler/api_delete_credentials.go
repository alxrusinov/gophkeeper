package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DeleteCredentials - deletes login and password data
func (h *HttpHandler) DeleteCredentials(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)

}
