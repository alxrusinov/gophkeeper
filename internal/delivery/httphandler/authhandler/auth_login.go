package authhandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func isNonExistUserError(err error) (res bool) {
	nonExistentErr := new(customerrors.NonexistentUser)

	if err != nil && errors.As(err, &nonExistentErr) {
		res = true
	}

	return res
}

func isUnverifiedUserError(err error) (res bool) {
	unverifiedErr := new(customerrors.UnverifiedUser)

	if err != nil && errors.As(err, &unverifiedErr) {
		res = true
	}
	return res
}

// Login - method of loginng existing user
func (ah *AuthHandler) Login(ctx iris.Context) {
	login := new(model.Login)
	err := ctx.ReadJSON(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	user, err := ah.usecase.VerifyUser(ctx, login)

	switch {
	case isNonExistUserError(err):
	case isUnverifiedUserError(err):
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	case err != nil:
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	tokenPair, err := ah.auth.GetTokenPair(user)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	tokenCookie := &http.Cookie{
		Name:     ah.userCookie,
		Value:    tokenPair.AccessToken,
		Expires:  time.Now().Add(ah.auth.GetAccessTokenExp()),
		Path:     "/",
		HttpOnly: true,
	}

	ctx.SetCookie(tokenCookie)

	ctx.StatusCode(http.StatusOK)

	ctx.JSON(user.ID)
}
