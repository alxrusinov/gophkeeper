package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetBinary - return binary data for user
func (u *Usecase) GetBinary(ctx context.Context, userID string, binID string) (*model.Binary, error) {
	return u.repository.GetBinary(ctx, userID, binID)
}
