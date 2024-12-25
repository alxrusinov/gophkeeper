package mongo

import (
	"context"
	"errors"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddCredentials - adds new credentials for user
func (m *Mongo) AddCredentials(ctx context.Context, creds *model.Credentials) (*model.Credentials, error) {
	insertedResult, err := m.client.Database(DataBase).Collection(CredentialsCollection).InsertOne(ctx, creds)

	if err != nil {
		return nil, err
	}

	if id, ok := insertedResult.InsertedID.(primitive.ObjectID); ok {
		return &model.Credentials{
			ID:     id.Hex(),
			UserID: "",
			Data:   creds.Data,
			Title:  creds.Title,
			Meta:   creds.Meta,
		}, nil
	}
	return nil, errors.New("credentials was not saved")
}
