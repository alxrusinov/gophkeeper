package notehandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/kataras/iris/v12"
)

func (nh *NoteHandler) GetNote(ctx iris.Context) {
	noteID := ctx.Params().Get("id")

	if noteID == "" {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	user, err := nh.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	note, err := nh.usecase.GetNote(ctx, user.ID, noteID)

	if err != nil {
		notFoundErr := new(customerrors.NotFound)

		if errors.As(err, &notFoundErr) {
			ctx.StopWithStatus(http.StatusNotFound)
			return
		}

		ctx.StopWithStatus(http.StatusInternalServerError)
		return

	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(note)
}
