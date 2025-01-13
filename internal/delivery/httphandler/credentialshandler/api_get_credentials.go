package credentialshandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/kataras/iris/v12"
)

func (ch *CredentialsHandler) GetCredentials(ctx iris.Context) {
	credID := ctx.Params().Get("id")

	if credID == "" {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	user, err := ch.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	creds, err := ch.usecase.GetCredentials(ctx, user.ID, credID)

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
	ctx.JSON(creds)
}
