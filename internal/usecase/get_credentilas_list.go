package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetCredentialsList - return all credentials for user
func (u *Usecase) GetCredentialsList(ctx context.Context, userID string) ([]model.Credentials, error) {
	return u.repository.GetCredentialsList(ctx, userID)
}
