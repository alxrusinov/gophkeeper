package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetBinaryList - return all binary data for user
func (u *Usecase) GetBinaryList(ctx context.Context, userID string) ([]model.Binary, error) {
	return u.repository.GetBinaryList(ctx, userID)
}
