package authhandler

import (
	"context"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// Auth - interface auth for handler
type Auth interface {
	// GetAccessTokenExp - return expiring of access token
	GetAccessTokenExp() time.Duration
	// GetAccessToken - generate access token
	GetAccessToken(user *model.User) (string, error)
	// GetAccessToken - generate access token
	GetRefreshToken(user *model.User) (string, error)
	// GetSigKey - returns sig aky  as []byte
	GetSigKey() []byte
	// GetTokenPair - return new token pair for user
	GetTokenPair(user *model.User) (*model.TokenPair, error)
	// GetVerifier - return custom jwt verifier
	GetVerifier() *jwt.Verifier
	// RefreshUserTokens refreshes access and refresh user tokens
	RefreshUserTokens(ctx iris.Context) (*model.TokenPair, error)
	// GetUserIDFromContext extracts userID from context
	GetUserFromContext(ctx iris.Context) (*model.User, error)
}

// Usecase - interface of Usecase
type Usecase interface {
	// VerifyUser - return information about user, if user exists
	// fact of user existing and error
	VerifyUser(ctx context.Context, lg *model.Login) (*model.User, error)
	// CreateUser - create new user
	CreateUser(ctx context.Context, lg *model.Login) (*model.User, error)
	// CheckUser - checks if user from token existss in repository
	CheckUser(ctx context.Context, userID string) (bool, error)
}

// AuthHandler - authorization handler
type AuthHandler struct {
	auth       Auth
	usecase    Usecase
	userCookie string
}

// NewAuthHandler - construct authorization handler
func NewAuthHandler(usecase Usecase, auth Auth, userCookie string) *AuthHandler {
	return &AuthHandler{
		auth:       auth,
		usecase:    usecase,
		userCookie: userCookie,
	}
}
