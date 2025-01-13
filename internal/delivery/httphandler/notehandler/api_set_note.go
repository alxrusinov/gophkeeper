package notehandler

import (
	"fmt"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (nh *NoteHandler) SetNote(ctx iris.Context) {
	user, err := nh.auth.GetUserFromContext(ctx)

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

	note.UserID = user.ID

	addedNote, err := nh.usecase.AddNote(ctx, note)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, fmt.Errorf("note with title %s was not saved", note.Title))
	}

	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(addedNote)
}
