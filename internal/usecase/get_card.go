package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// GetBankCard - return bank card for user
func (u *Usecase) GetBankCard(ctx context.Context, userID string, cardID string) (*model.BankCard, error) {
	return u.repository.GetBankCard(ctx, userID, cardID)
}
