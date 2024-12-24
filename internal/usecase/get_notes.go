package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetNote - returns note for user by note id
func (u *Usecase) GetNotes(ctx context.Context, userID string) ([]model.Note, error) {
	return u.repository.GetNotes(ctx, userID)
}
