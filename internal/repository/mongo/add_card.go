package mongo

import (
	"context"
	"errors"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddBankCard - adds new bank card for user
func (m *Mongo) AddBankCard(ctx context.Context, card *model.BankCard) (*model.BankCard, error) {

	insertedResult, err := m.client.Database(DataBase).Collection(BankCardCollection).InsertOne(ctx, card)

	if err != nil {
		return nil, err
	}

	if id, ok := insertedResult.InsertedID.(primitive.ObjectID); ok {
		return &model.BankCard{
			ID:     id.Hex(),
			UserID: "",
			Data:   card.Data,
			Title:  card.Title,
			Meta:   card.Meta,
		}, nil
	}
	return nil, errors.New("credentials was not saved")
}
