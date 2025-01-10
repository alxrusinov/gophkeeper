package auth

import (
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// RefreshUserTokens refreshes access and refresh user tokens
func (a *Auth) RefreshUserTokens(ctx iris.Context) (*model.TokenPair, error) {
	user, err := a.GetUserFromContext(ctx)

	accessToken, err := a.GetAccessToken(user)

	if err != nil {
		return nil, err
	}

	refreshToken, err := a.GetRefreshToken(user)

	if err != nil {
		return nil, err
	}

	tokenPair := &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Exp:          int64(a.config.Auth.AccessExpire.Seconds()),
	}

	return tokenPair, nil

}
