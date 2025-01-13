package authhandler

import (
	"net/http"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// Register- method of registration new user
func (ah *AuthHandler) Register(ctx iris.Context) {
	login := new(model.Login)

	err := ctx.ReadJSON(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	user, err := ah.usecase.CreateUser(ctx, login)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	accessToken, err := ah.auth.GetAccessToken(user)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     ah.userCookie,
		Value:    accessToken,
		Expires:  time.Now().Add(ah.auth.GetAccessTokenExp()),
		Path:     "/",
		HttpOnly: true,
	}

	ctx.SetCookie(cookie)
	ctx.StatusCode(http.StatusCreated)
}
