package httphandler

import (
	"time"

	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler/authhandler"
	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler/bankcardhandler"
	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler/binaryhandler"
	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler/credentialshandler"
	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler/middleware"
	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler/notehandler"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// HttpHandler - general handler included other handlers
type HttpHandler struct {
	AuthHandler        *authhandler.AuthHandler
	NoteHandler        *notehandler.NoteHandler
	CredentialsHandler *credentialshandler.CredentialsHandler
	BinaryHandler      *binaryhandler.BinaryHandler
	BankCardHandler    *bankcardhandler.BankCardHandler
	Middleware         *middleware.Middleware
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

// Usecase - interface of Usecase
type Usecase interface {
	notehandler.Usecase
	credentialshandler.Usecase
	binaryhandler.Usecase
	bankcardhandler.Usecase
	authhandler.Usecase
}

// New HttpHandler - return new instsce of HttpHandler
func NewHttpHandler(usecase Usecase, auth Auth) *HttpHandler {
	return &HttpHandler{
		AuthHandler:        authhandler.NewAuthHandler(usecase, auth, userCookie),
		NoteHandler:        notehandler.NewNoteHandlerHandler(usecase, auth),
		CredentialsHandler: credentialshandler.NewCredentialsHandler(usecase, auth),
		BankCardHandler:    bankcardhandler.NewBankCardHandler(usecase, auth),
		BinaryHandler:      binaryhandler.NewBinaryHandlerHandler(usecase, auth),

		Middleware: middleware.NewMiddleware(usecase, auth, userCookie),
	}
}
