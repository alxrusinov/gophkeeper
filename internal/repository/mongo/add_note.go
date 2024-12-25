package mongo

import (
	"context"
	"errors"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddNote - adds new note for user
func (m *Mongo) AddNote(ctx context.Context, note *model.Note, userID string) (*model.Note, error) {
	insertedResult, err := m.client.Database(DataBase).Collection(NoteCollection).InsertOne(ctx, note)

	if err != nil {
		return nil, err
	}

	if id, ok := insertedResult.InsertedID.(primitive.ObjectID); ok {
		return &model.Note{
			ID:     id.Hex(),
			UserID: note.UserID,
			Data:   note.Data,
			Title:  note.Title,
			Meta:   note.Meta,
		}, nil
	}

	return nil, errors.New("note was not saved")
}
