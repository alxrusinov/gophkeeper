package mongo

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetCredentialsList - return all credentials for user
func (m *Mongo) GetCredentialsList(ctx context.Context, userID string) ([]model.Credentials, error) {
	result := make([]model.Credentials, 0)

	filter := bson.D{{Key: "user_id", Value: userID}}

	cursor, err := m.client.Database(DataBase).Collection(CredentialsCollection).Find(ctx, filter, options.Find().SetProjection(bson.M{"user_id": 0}))

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		item := new(model.CredentialsDocument)

		err := cursor.Decode(item)

		if err != nil {
			return nil, err
		}

		note := model.CredentialsFromCredentialsDocument(*item)

		result = append(result, *note)
	}

	return result, nil
}
