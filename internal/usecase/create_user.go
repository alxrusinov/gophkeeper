package usecase

import "github.com/alxrusinov/gophkeeper/internal/model"

// CreateUser - create new user
func (u *Usecase) CreateUser(lg *model.Login) (*model.User, error) {
	return &model.User{
		ID:       "123",
		Username: lg.Username,
	}, nil
}
