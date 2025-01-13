package notehandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (nh *NoteHandler) GetNoteList(ctx iris.Context) {
	user, err := nh.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	notes, err := nh.usecase.GetNoteList(ctx, user.ID)

	if err != nil {
		notFoundErr := new(customerrors.NotFound)

		if errors.As(err, &notFoundErr) {
			ctx.StatusCode(http.StatusNotFound)
			ctx.JSON([]model.Note{})
			return
		}

		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(notes)
}
