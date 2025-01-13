package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetNote - returns note for user by note id
func (u *Usecase) GetNote(ctx context.Context, userID, noteID string) (*model.Note, error) {
	return u.repository.GetNote(ctx, userID, noteID)
}
