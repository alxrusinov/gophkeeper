package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CheckUser - checks if user from token existss in repository
func (m *Mongo) CheckUser(ctx context.Context, userID string) (bool, error) {
	var user model.UserDocument

	filterUserID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return false, err
	}

	filter := bson.D{{Key: "_id", Value: filterUserID}}

	result := m.client.Database(DataBase).Collection(UserCollection).FindOne(ctx, filter)

	err = result.Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, &customerrors.NonexistentUser{Err: fmt.Errorf("user winth %s does not exist", userID)}
		}
		return false, err
	}

	return true, nil

}
