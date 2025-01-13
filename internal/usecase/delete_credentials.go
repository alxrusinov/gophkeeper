package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// DeleteBankCard - delete credentials
func (u *Usecase) DeleteCredentials(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	return u.repository.DeleteCredentials(ctx, source)
}
