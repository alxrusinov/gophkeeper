package httphandler

import (
	"net/http"

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

	note, err := h.usecase.GetNote(user.ID, noteID)

	if err != nil {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	note.UserID = ""

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(note)
}
