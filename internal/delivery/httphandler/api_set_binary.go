package httphandler

import (
	"fmt"
	"net/http"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) SetBinary(ctx iris.Context) {
	user, err := h.auth.GetUserFromContext(ctx)

	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	binary := new(model.Binary)

	err = ctx.ReadJSON(binary)

	if err != nil {
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}

	binary.UserID = user.ID

	added, err := h.usecase.AddBinary(ctx, binary)

	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, fmt.Errorf("credentils with title %s was not saved", added.Title))
	}

	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(added)
}
