package usecase

import "github.com/alxrusinov/gophkeeper/internal/model"

// AddNote - adds new note for user
func (u *Usecase) AddNote(note *model.Note, userID string) (*model.Note, error) {
	return u.repository.AddNote(note, userID)
}
