package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// AddBinary - adds new binaey data for user
func (u *Usecase) AddBinary(ctx context.Context, data *model.BinaryUpload) (*model.Binary, error) {
	return u.repository.AddBinary(ctx, data)
}
