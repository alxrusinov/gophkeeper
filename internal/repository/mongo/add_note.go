package mongo

import "github.com/alxrusinov/gophkeeper/internal/model"

// AddNote - adds new note for user
func (m *Mongo) AddNote(note *model.Note, userID string) (*model.Note, error) {
	return note, nil
}
