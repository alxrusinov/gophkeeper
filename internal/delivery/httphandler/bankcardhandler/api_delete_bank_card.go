package bankcardhandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// DeleteBankCard - deletes bank card
func (bk *BankCardHandler) DeleteBankCard(ctx iris.Context) {
	user, err := bk.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	cardID := new(model.SourceID)

	err = ctx.ReadJSON(cardID)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	cardID.UserID = user.ID

	res, err := bk.usecase.DeleteBankCard(ctx, cardID)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, errors.New("note was not deleted"))
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(res)
}
