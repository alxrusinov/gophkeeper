package mock

import (
	"time"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/stretchr/testify/mock"
)

const (
	indexZero int = iota
	indexFirst
	indexSecond
)

// AuthMock - mocked auth structure
type AuthMock struct {
	mock.Mock
}

// GetAccessTokenExp - return expiring of access token
func (am *AuthMock) GetAccessTokenExp() time.Duration {
	args := am.Called()

	return args.Get(indexZero).(time.Duration)
}

// GetAccessToken - generate access token
func (am *AuthMock) GetAccessToken(user *model.User) (string, error) {
	args := am.Called(user)

	return args.String(indexZero), args.Error(indexFirst)
}

// GetAccessToken - generate access token
func (am *AuthMock) GetRefreshToken(user *model.User) (string, error) {
	args := am.Called(user)

	return args.String(indexZero), args.Error(indexFirst)
}

// GetSigKey - returns sig aky  as []byte
func (am *AuthMock) GetSigKey() []byte {
	args := am.Called()

	return args.Get(indexZero).([]byte)
}

// GetTokenPair - return new token pair for user
func (am *AuthMock) GetTokenPair(user *model.User) (*model.TokenPair, error) {
	args := am.Called(user)

	return args.Get(indexZero).(*model.TokenPair), args.Error(indexFirst)
}

// GetVerifier - return custom jwt verifier
func (am *AuthMock) GetVerifier() *jwt.Verifier {
	args := am.Called()

	return args.Get(indexZero).(*jwt.Verifier)
}

// RefreshUserTokens refreshes access and refresh user tokens
func (am *AuthMock) RefreshUserTokens(ctx iris.Context) (*model.TokenPair, error) {
	args := am.Called(ctx)

	return args.Get(indexZero).(*model.TokenPair), args.Error(indexFirst)
}

// GetUserIDFromContext extracts userID from context
func (am *AuthMock) GetUserFromContext(ctx iris.Context) (*model.User, error) {
	args := am.Called(ctx)

	return args.Get(indexZero).(*model.User), args.Error(indexFirst)
}

// NewAuthMock - mocked auth structure
func NewAuthMock() *AuthMock {
	return new(AuthMock)
}
