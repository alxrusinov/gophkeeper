package bankcardhandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/kataras/iris/v12"
)

func (bk *BankCardHandler) GetBankCard(ctx iris.Context) {
	cardID := ctx.Params().Get("id")

	if cardID == "" {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	user, err := bk.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	card, err := bk.usecase.GetBankCard(ctx, user.ID, cardID)

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
	ctx.JSON(card)
}
