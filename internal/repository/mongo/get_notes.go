package mongo

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetNotes - return notes for user by note id
func (m *Mongo) GetNotes(ctx context.Context, userID string) ([]model.Note, error) {

	result := make([]model.Note, 0)

	filter := bson.D{{Key: "user_id", Value: userID}}

	cursor, err := m.client.Database(DataBase).Collection(NoteCollection).Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		item := model.Note{}

		err := cursor.Decode(&item)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil

}
