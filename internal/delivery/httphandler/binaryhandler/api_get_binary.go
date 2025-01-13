package binaryhandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/kataras/iris/v12"
)

func (b *BinaryHandler) GetBinary(ctx iris.Context) {
	binID := ctx.Params().Get("id")

	if binID == "" {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	user, err := b.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	bin, err := b.usecase.GetBinary(ctx, user.ID, binID)

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
	ctx.JSON(bin)
}
