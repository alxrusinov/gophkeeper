package auth

import "github.com/alxrusinov/gophkeeper/internal/model"

// GetTokenPair - return new token pair for user
func (a *Auth) GetTokenPair(user *model.User) (*model.TokenPair, error) {
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
		Exp:          int64(accessExpire.Seconds()),
	}

	return tokenPair, nil
}
