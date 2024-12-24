package mongo

import (
	"context"
	"errors"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser - create new user
func (m *Mongo) CreateUser(lg *model.Login) (*model.User, error) {
	users := m.client.Database(DataBase).Collection(UserCollection)

	res, err := users.InsertOne(context.Background(), lg)

	if err != nil {
		return nil, err
	}

	if userID, ok := res.InsertedID.(primitive.ObjectID); ok {
		return &model.User{
			ID:       userID.String(),
			Username: lg.Username,
		}, nil
	}

	return nil, errors.New("user was not created")

}
