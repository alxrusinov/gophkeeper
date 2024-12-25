package mongo

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetNoteList - return notes for user by note id
func (m *Mongo) GetNoteList(ctx context.Context, userID string) ([]model.Note, error) {

	result := make([]model.Note, 0)

	filter := bson.D{{Key: "user_id", Value: userID}}

	cursor, err := m.client.Database(DataBase).Collection(NoteCollection).Find(ctx, filter, options.Find().SetProjection(bson.M{"user_id": 0}))

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		item := new(model.NoteDocument)

		err := cursor.Decode(item)

		if err != nil {
			return nil, err
		}

		note := model.NoteFromNoteDocument(*item)

		result = append(result, *note)
	}

	return result, nil

}
