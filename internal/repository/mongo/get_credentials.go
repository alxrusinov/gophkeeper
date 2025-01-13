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

// GetCredentials - return credentials for user
func (m *Mongo) GetCredentials(ctx context.Context, userID string, credsID string) (*model.Credentials, error) {
	result := new(model.CredentialsDocument)

	credsIDFilter, err := primitive.ObjectIDFromHex(credsID)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("credentials with id %s does not exist", credsID)}
	}

	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: credsIDFilter}}

	err = m.client.Database(DataBase).Collection(CredentialsCollection).FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"user_id": 0})).Decode(result)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("note with id %s does not exist", credsID)}
	}

	credentials := model.CredentialsFromCredentialsDocument(*result)

	return credentials, nil
}
