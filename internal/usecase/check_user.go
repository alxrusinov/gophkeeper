package usecase

import "context"

// CheckUser - checks if user from token existss in repository
func (u *Usecase) CheckUser(ctx context.Context, userID string) (bool, error) {
	return u.repository.CheckUser(ctx, userID)
}
