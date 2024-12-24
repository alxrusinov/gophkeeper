package mongo

import (
	"context"
	"fmt"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetNote - returns note for user by note id
func (m *Mongo) GetNote(ctx context.Context, userID, noteID string) (*model.Note, error) {
	result := new(model.Note)

	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: noteID}}

	err := m.client.Database(DataBase).Collection(NoteCollection).FindOne(ctx, filter).Decode(result)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("note with id %s does not exist", noteID)}
	}

	return result, nil
}
