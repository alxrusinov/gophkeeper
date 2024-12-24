package httphandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// Login - method of loginng existing user
func (h *HttpHandler) Login(ctx iris.Context) {
	login := new(model.Login)
	err := ctx.ReadJSON(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := h.usecase.VerifyUser(login)

	if err != nil {
		nonExistentErr := new(customerrors.NonexistentUser)

		if ok := errors.As(err, &nonExistentErr); ok {
			ctx.StopWithStatus(http.StatusUnauthorized)
			return
		}

		unverifiedErr := new(customerrors.UnverifiedUser)

		if ok := errors.As(err, &unverifiedErr); ok {
			ctx.StopWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	tokenPair, err := h.auth.GetTokenPair(user)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	tokenCookie := &http.Cookie{
		Name:     userCookie,
		Value:    tokenPair.AccessToken,
		Expires:  time.Now().Add(h.auth.GetAccessTokenExp()),
		Path:     "/",
		HttpOnly: true,
	}

	ctx.SetCookie(tokenCookie)

	ctx.StatusCode(http.StatusOK)
}
