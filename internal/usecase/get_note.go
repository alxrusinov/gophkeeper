package usecase

import "github.com/alxrusinov/gophkeeper/internal/model"

// GetNote - returns note for user by note id
func (u *Usecase) GetNote(userID, noteID string) (*model.Note, error) {
	return u.repository.GetNote(userID, noteID)
}
