package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// VerifyUser - checks if user exists and has valid password
func (m *Mongo) VerifyUser(lg *model.Login) (*model.User, error) {
	user := new(model.User)
	login := new(model.Login)

	result := m.client.Database(DataBase).Collection(UserCollection).FindOne(context.Background(), bson.M{"username": lg.Username})

	err := result.Decode(user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, &customerrors.NonexistentUser{Err: fmt.Errorf("user winth %s does not exist", lg.Username)}
		}
		return nil, err
	}

	err = result.Decode(login)

	if err != nil {
		return nil, err
	}

	if lg.Password != login.Password {
		return nil, &customerrors.UnverifiedUser{Err: fmt.Errorf("login or password is wrong")}
	}

	return user, nil
}
