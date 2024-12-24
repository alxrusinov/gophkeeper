package usecase

import "github.com/alxrusinov/gophkeeper/internal/model"

// CreateUser - create new user
func (u *Usecase) CreateUser(lg *model.Login) (*model.User, error) {
	return u.repository.CreateUser(lg)
}
