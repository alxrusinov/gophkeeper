package middleware

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// Middlware - structure of middlware
type Middleware struct {
	auth       Auth
	usecase    Usecase
	userCookie string
}

// Auth - authorization
type Auth interface {
	// GetVerifier - return custom jwt verifier
	GetVerifier() *jwt.Verifier
	// GetUserIDFromContext extracts userID from context
	GetUserFromContext(ctx iris.Context) (*model.User, error)
}

// Usecase - interface of Usecase
type Usecase interface {
	// CheckUser - checks if user from token existss in repository
	CheckUser(ctx context.Context, userID string) (bool, error)
}

// NewMiddleware - create new instance of Middleware
func NewMiddleware(usecase Usecase, auth Auth, userCookie string) *Middleware {
	return &Middleware{
		auth:       auth,
		userCookie: userCookie,
		usecase:    usecase,
	}
}
