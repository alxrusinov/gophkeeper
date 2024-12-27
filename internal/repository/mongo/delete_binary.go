package mongo

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteBankCard - delete binary
func (m *Mongo) DeleteBinary(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	sourceFilter, err := primitive.ObjectIDFromHex(source.ID)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "user_id", Value: source.UserID}, {Key: "_id", Value: sourceFilter}}

	_, err = m.client.Database(DataBase).Collection(BinaryCollection).DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return source, nil
}
