package mongo

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetBinaryList - return all binary data for user
func (m *Mongo) GetBinaryList(ctx context.Context, userID string) ([]model.Binary, error) {
	result := make([]model.Binary, 0)

	filter := bson.D{{Key: "user_id", Value: userID}}

	cursor, err := m.client.Database(DataBase).Collection(BinaryCollection).Find(ctx, filter, options.Find().SetProjection(bson.M{"user_id": 0}))

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		item := new(model.BinaryDocument)

		err := cursor.Decode(item)

		if err != nil {
			return nil, err
		}

		binData := model.BinaryFromBinaryDocument(*item)

		result = append(result, *binData)
	}

	return result, nil
}
