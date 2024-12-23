package router

import (
	"net/http"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// Register- method of registration new user
func (r *Router) Register(ctx iris.Context) {
	login := new(model.Login)

	err := ctx.ReadJSON(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	user, err := r.usecase.CreateUser(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	accessToken, err := r.auth.GetAccessToken(user)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:    userCookie,
		Value:   accessToken,
		Expires: time.Now().Add(r.auth.GetAccessTokenExp()),
		Path:    apiRouteGroup,
	}

	ctx.SetCookie(cookie)
	ctx.StatusCode(http.StatusCreated)
}
