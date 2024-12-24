package httphandler

import (
	"fmt"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetNote(ctx iris.Context) {
	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	note := new(model.Note)

	err = ctx.ReadJSON(note)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	addedNote, err := h.usecase.AddNote(note, user.ID)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, fmt.Errorf("note with title %s was not saved", note.Title))
	}

	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(addedNote)
}
