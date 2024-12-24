package usecase

import "github.com/alxrusinov/gophkeeper/internal/model"

// VerifyUser - checks if user exists and has valid password
func (u *Usecase) VerifyUser(lg *model.Login) (*model.User, error) {
	return u.repository.VerifyUser(lg)
}
