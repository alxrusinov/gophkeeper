package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// VerifyUser - checks if user exists and has valid password
func (u *Usecase) VerifyUser(ctx context.Context, lg *model.Login) (*model.User, error) {
	return u.repository.VerifyUser(ctx, lg)
}
