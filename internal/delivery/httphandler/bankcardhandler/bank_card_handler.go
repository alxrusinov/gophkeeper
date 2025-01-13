package bankcardhandler

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// BankCard - bankcard handler
type BankCardHandler struct {
	auth    Auth
	usecase Usecase
}

// Auth - interface auth for handler
type Auth interface {
	// GetUserIDFromContext extracts userID from context
	GetUserFromContext(ctx iris.Context) (*model.User, error)
}

// Usecase - interface of Usecase
type Usecase interface {
	// AddBankCard - adds new bank card for user
	AddBankCard(ctx context.Context, card *model.BankCard) (*model.BankCard, error)
	// GetBankCard - return bank card for user
	GetBankCard(ctx context.Context, userID string, cardID string) (*model.BankCard, error)
	// GetBankCardList - return all bank cards for user
	GetBankCardList(ctx context.Context, userID string) ([]model.BankCard, error)
	// DeleteBankCard - delete bank card
	DeleteBankCard(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
}

// NewBankCardHandler - crete new instance of bank card structure
func NewBankCardHandler(usecase Usecase, auth Auth) *BankCardHandler {
	return &BankCardHandler{
		auth:    auth,
		usecase: usecase,
	}
}
