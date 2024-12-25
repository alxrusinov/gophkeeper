package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetCredentials - return credentials for user
func (u *Usecase) GetCredentials(ctx context.Context, userID string, credsID string) (*model.Credentials, error) {

	return u.repository.GetCredentials(ctx, userID, credsID)
}
