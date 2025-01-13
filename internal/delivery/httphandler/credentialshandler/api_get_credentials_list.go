package credentialshandler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (ch *CredentialsHandler) GetCredentialsList(ctx iris.Context) {

	user, err := ch.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	credsList, err := ch.usecase.GetCredentialsList(ctx, user.ID)

	if err != nil {
		notFoundErr := new(customerrors.NotFound)

		if errors.As(err, &notFoundErr) {
			ctx.StatusCode(http.StatusNotFound)
			ctx.JSON([]model.Credentials{})
			return
		}

		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(credsList)
}
