package httphandler

import (
	"net/http"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// Register- method of registration new user
func (h *HttpHandler) Register(ctx iris.Context) {
	login := new(model.Login)

	err := ctx.ReadJSON(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	user, err := h.usecase.CreateUser(login)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	accessToken, err := h.auth.GetAccessToken(user)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     userCookie,
		Value:    accessToken,
		Expires:  time.Now().Add(h.auth.GetAccessTokenExp()),
		Path:     ApiRouteGroup,
		HttpOnly: true,
	}

	ctx.SetCookie(cookie)
	ctx.StatusCode(http.StatusCreated)
}
