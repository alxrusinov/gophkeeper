package mongo

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetBankCardList - return all bank cards for user
func (m *Mongo) GetBankCardList(ctx context.Context, userID string) ([]model.BankCard, error) {
	result := make([]model.BankCard, 0)

	filter := bson.D{{Key: "user_id", Value: userID}}

	cursor, err := m.client.Database(DataBase).Collection(BankCardCollection).Find(ctx, filter, options.Find().SetProjection(bson.M{"user_id": 0}))

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		item := new(model.BankCardDocument)

		err := cursor.Decode(item)

		if err != nil {
			return nil, err
		}

		card := model.BankCardFromBankCardDocument(*item)

		result = append(result, *card)
	}

	return result, nil
}
