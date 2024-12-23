package httphandler

import (
	"github.com/alxrusinov/gophkeeper/internal/auth"
	"github.com/alxrusinov/gophkeeper/internal/model"
)

// Usecase - interface of Usecase
type Usecase interface {
	// VerifyUser - return information about user, if user exists
	// fact of user existing and error
	VerifyUser(lg *model.Login) (*model.User, error)
	// CreateUser - create new user
	CreateUser(lg *model.Login) (*model.User, error)
}

// HttpHandler - handler for http router
type HttpHandler struct {
	auth    *auth.Auth
	usecase Usecase
}

// NewHttpHandler - return new instance
func NewHttpHandler(usecase Usecase) *HttpHandler {
	return &HttpHandler{
		auth:    auth.NewAuth(),
		usecase: usecase,
	}
}
