package credentialshandler

import (
	"context"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// CredentialsHandler - handler for credentials data
type CredentialsHandler struct {
	usecase Usecase
	auth    Auth
}

// Usecsae - interface for usecase
type Usecase interface {
	// AddCredentials - adds new credentials for user
	AddCredentials(ctx context.Context, creds *model.Credentials) (*model.Credentials, error)
	// GetCredentials - return credentials for user
	GetCredentials(ctx context.Context, userID string, credsID string) (*model.Credentials, error)
	// GetCredentialsList - return all credentials for user
	GetCredentialsList(ctx context.Context, userID string) ([]model.Credentials, error)
	// DeleteBankCard - delete credentials
	DeleteCredentials(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
}

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

// NewCredentialsHandler - returns new CredentialsHandler
func NewCredentialsHandler(usecase Usecase, auth Auth) *CredentialsHandler {
	return &CredentialsHandler{
		auth:    auth,
		usecase: usecase,
	}
}
