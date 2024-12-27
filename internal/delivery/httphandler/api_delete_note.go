package httphandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) DeleteNote(ctx iris.Context) {
	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	noteID := new(model.SourceID)

	err = ctx.ReadJSON(noteID)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	noteID.UserID = user.ID

	res, err := h.usecase.DeleteNote(ctx, noteID)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, errors.New("note was not deleted"))
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(res)
}
