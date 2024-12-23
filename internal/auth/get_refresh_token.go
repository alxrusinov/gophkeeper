package auth

import (
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// GetAccessToken - generate access token
func (a *Auth) GetRefreshToken(user *model.User) (string, error) {
	refreshClaims := Claims{
		Claims: jwt.Claims{Subject: user.Username},
		UserID: user.ID,
	}

	refreshToken, err := a.refreshSigner.Sign(refreshClaims)

	return string(refreshToken), err
}
