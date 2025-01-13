package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// CreateUser - create new user
func (u *Usecase) CreateUser(ctx context.Context, lg *model.Login) (*model.User, error) {
	return u.repository.CreateUser(ctx, lg)
}
