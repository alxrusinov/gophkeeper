package httphandler

import (
	"fmt"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetBankCard(ctx iris.Context) {
	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	card := new(model.BankCard)

	err = ctx.ReadJSON(card)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	card.UserID = user.ID

	addedCard, err := h.usecase.AddBankCard(ctx, card)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, fmt.Errorf("card with title %s was not saved", card.Title))
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(addedCard)
}
