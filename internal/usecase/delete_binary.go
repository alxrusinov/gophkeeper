package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// DeleteBankCard - delete binary
func (u *Usecase) DeleteBinary(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	return u.repository.DeleteBinary(ctx, source)
}
