package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// DeleteBankCard - delete note
func (u *Usecase) DeleteNote(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	return u.repository.DeleteNote(ctx, source)
}
