package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetBankCardList - return all bank cards for user
func (u *Usecase) GetBankCardList(ctx context.Context, userID string) ([]model.BankCard, error) {
	return u.repository.GetBankCardList(ctx, userID)
}
