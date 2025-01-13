package credentialshandler

import (
	"fmt"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (ch *CredentialsHandler) SetCredentials(ctx iris.Context) {
	user, err := ch.auth.GetUserFromContext(ctx)

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

	creds.UserID = user.ID

	addedCreds, err := ch.usecase.AddCredentials(ctx, creds)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, fmt.Errorf("credentils with title %s was not saved", creds.Title))
	}

	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(addedCreds)
}
