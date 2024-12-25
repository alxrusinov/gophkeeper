package mongo

import (
	"context"
	"fmt"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetBankCard - return bank card for user
func (m *Mongo) GetBankCard(ctx context.Context, userID string, cardID string) (*model.BankCard, error) {
	result := new(model.BankCardDocument)

	cardFilter, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("credentials with id %s does not exist", cardID)}
	}

	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: cardFilter}}

	err = m.client.Database(DataBase).Collection(BankCardCollection).FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"user_id": 0})).Decode(result)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("note with id %s does not exist", cardID)}
	}

	card := model.BankCardFromBankCardDocument(*result)

	return card, nil

}
