package router

import (
	"errors"
	"net/http"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// Login - method of loginng existing user
func (r *Router) Login(ctx iris.Context) {
	login := new(model.Login)
	err := ctx.ReadJSON(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := r.usecase.VerifyUser(login)

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

	tokenPair, err := r.auth.GetTokenPair(user)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	tokenCookie := &http.Cookie{
		Name:    userCookie,
		Value:   tokenPair.AccessToken,
		Expires: time.Now().Add(r.auth.GetAccessTokenExp()),
		Path:    apiRouteGroup,
	}

	ctx.SetCookie(tokenCookie)

	ctx.StatusCode(http.StatusOK)
}
