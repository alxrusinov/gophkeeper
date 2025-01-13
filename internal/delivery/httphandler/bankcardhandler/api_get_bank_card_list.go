package bankcardhandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (bk *BankCardHandler) GetBankCardList(ctx iris.Context) {

	user, err := bk.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	cardList, err := bk.usecase.GetBankCardList(ctx, user.ID)

	if err != nil {
		notFoundErr := new(customerrors.NotFound)

		if errors.As(err, &notFoundErr) {
			ctx.StatusCode(http.StatusNotFound)
			ctx.JSON([]model.BankCard{})
			return
		}

		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(cardList)
}
