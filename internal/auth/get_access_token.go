package auth

import (
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// GetAccessToken - generate access token
func (a *Auth) GetAccessToken(user *model.User) (string, error) {
	accessClaims := Claims{
		Claims: jwt.Claims{Subject: user.Username},
		UserID: user.ID,
	}

	accessToken, err := a.accessSigner.Sign(accessClaims)

	return string(accessToken), err
}
