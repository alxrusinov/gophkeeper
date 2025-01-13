package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// AddNote - adds new note for user
func (u *Usecase) AddNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	return u.repository.AddNote(ctx, note)
}
