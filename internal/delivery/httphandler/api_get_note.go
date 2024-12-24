package httphandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) GetNote(ctx iris.Context) {
	noteID := ctx.Params().Get("id")

	if noteID == "" {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	note, err := h.usecase.GetNote(ctx, user.ID, noteID)

	if err != nil {
		notFoundErr := new(customerrors.NotFound)

		if errors.As(err, &notFoundErr) {
			ctx.StopWithStatus(http.StatusNotFound)
			return
		}

		ctx.StopWithStatus(http.StatusInternalServerError)
		return

	}

	note.UserID = ""

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(note)
}
