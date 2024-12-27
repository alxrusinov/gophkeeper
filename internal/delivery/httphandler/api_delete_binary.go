package httphandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// DeleteBinary - deletes binary data
func (h *HttpHandler) DeleteBinary(ctx iris.Context) {
	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	binID := new(model.SourceID)

	err = ctx.ReadJSON(binID)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	binID.UserID = user.ID

	res, err := h.usecase.DeleteBinary(ctx, binID)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, errors.New("note was not deleted"))
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(res)
}
