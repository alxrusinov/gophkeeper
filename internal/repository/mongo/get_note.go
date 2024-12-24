package mongo

import "github.com/alxrusinov/gophkeeper/internal/model"

// GetNote - returns note for user by note id
func (m *Mongo) GetNote(userID, noteID string) (*model.Note, error) {
	return &model.Note{
		ID:     "111",
		UserID: "222",
		Title:  "Crazy title",
		Data:   "Some data",
		Meta:   "Some meta data",
	}, nil
}
