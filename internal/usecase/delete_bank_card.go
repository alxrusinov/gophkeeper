package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// DeleteBankCard - delete bank card
func (u *Usecase) DeleteBankCard(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	return u.repository.DeleteBankCard(ctx, source)
}
