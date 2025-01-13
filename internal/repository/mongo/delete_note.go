package mongo

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteBankCard - delete note
func (m *Mongo) DeleteNote(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	sourceFilter, err := primitive.ObjectIDFromHex(source.ID)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "user_id", Value: source.UserID}, {Key: "_id", Value: sourceFilter}}

	_, err = m.client.Database(DataBase).Collection(NoteCollection).DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return source, nil
}
