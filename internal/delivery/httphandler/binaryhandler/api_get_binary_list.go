package binaryhandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (b *BinaryHandler) GetBinaryList(ctx iris.Context) {

	user, err := b.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	binList, err := b.usecase.GetBinaryList(ctx, user.ID)

	if err != nil {
		notFoundErr := new(customerrors.NotFound)

		if errors.As(err, &notFoundErr) {
			ctx.StatusCode(http.StatusNotFound)
			ctx.JSON([]model.Binary{})
			return
		}

		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(binList)
}
