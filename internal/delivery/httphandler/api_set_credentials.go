package httphandler

import (
	"fmt"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetCredentials(ctx iris.Context) {
	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	creds := new(model.Credentials)

	err = ctx.ReadJSON(creds)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	addedCreds, err := h.usecase.AddCredentials(ctx, creds, user.ID)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, fmt.Errorf("credentils with title %s was not saved", creds.Title))
	}

	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(addedCreds)
}
