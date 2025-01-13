package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// AddCredentials - adds new credentials for user
func (u *Usecase) AddCredentials(ctx context.Context, creds *model.Credentials) (*model.Credentials, error) {
	return u.repository.AddCredentials(ctx, creds)
}
