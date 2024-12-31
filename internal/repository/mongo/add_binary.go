package mongo

import (
	"context"
	"errors"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddBinary - adds new binaey data for user
func (m *Mongo) AddBinary(ctx context.Context, data *model.Binary) (*model.Binary, error) {

	insertedResult, err := m.client.Database(DataBase).Collection(BinaryCollection).InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	if id, ok := insertedResult.InsertedID.(primitive.ObjectID); ok {
		return &model.Binary{
			ID:       id.Hex(),
			UserID:   "",
			Data:     data.Data,
			Title:    data.Title,
			Meta:     data.Meta,
			MimeType: data.MimeType,
		}, nil
	}
	return nil, errors.New("credentials was not saved")
}
