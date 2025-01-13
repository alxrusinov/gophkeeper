package auth

import (
	"errors"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// GetUserIDFromContext extracts userID from context
func (a *Auth) GetUserFromContext(ctx iris.Context) (*model.User, error) {
	token := jwt.GetVerifiedToken(ctx)

	if token == nil {
		return nil, errors.New("token not found")
	}

	var claims Claims
	if err := token.Claims(&claims); err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       claims.UserID,
		Username: claims.Subject,
	}

	return user, nil
}
