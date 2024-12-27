package httphandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// DeleteCredentials - deletes login and password data
func (h *HttpHandler) DeleteCredentials(ctx iris.Context) {
	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	credsID := new(model.SourceID)

	err = ctx.ReadJSON(credsID)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	credsID.UserID = user.ID

	res, err := h.usecase.DeleteCredentials(ctx, credsID)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, errors.New("note was not deleted"))
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(res)

}
