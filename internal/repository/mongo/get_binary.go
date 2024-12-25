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

// GetBinary - return binary data for user
func (m *Mongo) GetBinary(ctx context.Context, userID string, binID string) (*model.Binary, error) {
	result := new(model.BinaryDocument)

	binFilter, err := primitive.ObjectIDFromHex(binID)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("data with id %s does not exist", binID)}
	}

	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: binFilter}}

	err = m.client.Database(DataBase).Collection(BinaryCollection).FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"user_id": 0})).Decode(result)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("data with id %s does not exist", binID)}
	}

	binData := model.BinaryFromBinaryDocument(*result)

	return binData, nil
}
