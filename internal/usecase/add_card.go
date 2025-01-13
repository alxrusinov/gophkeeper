package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// AddBankCard - adds new bank card for user
func (u *Usecase) AddBankCard(ctx context.Context, card *model.BankCard) (*model.BankCard, error) {
	return u.repository.AddBankCard(ctx, card)
}
