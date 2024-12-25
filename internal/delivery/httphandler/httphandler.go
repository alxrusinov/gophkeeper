package httphandler

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/auth"
	"github.com/alxrusinov/gophkeeper/internal/model"
)

// Usecase - interface of Usecase
type Usecase interface {
	// VerifyUser - return information about user, if user exists
	// fact of user existing and error
	VerifyUser(ctx context.Context, lg *model.Login) (*model.User, error)
	// CreateUser - create new user
	CreateUser(ctx context.Context, lg *model.Login) (*model.User, error)
	// GetNote - return note for user by note id
	GetNote(ctx context.Context, userID string, noteID string) (*model.Note, error)
	// GetNotes - return notes for user by note id
	GetNoteList(ctx context.Context, userID string) ([]model.Note, error)
	// AddNote - adds new note for user
	AddNote(ctx context.Context, note *model.Note) (*model.Note, error)
	// AddCredentials - adds new credentials for user
	AddCredentials(ctx context.Context, creds *model.Credentials) (*model.Credentials, error)
	// GetCredentials - return credentials for user
	GetCredentials(ctx context.Context, userID string, credsID string) (*model.Credentials, error)
	// GetCredentialsList - return all credentials for user
	GetCredentialsList(ctx context.Context, userID string) ([]model.Credentials, error)
	// AddBankCard - adds new bank card for user
	AddBankCard(ctx context.Context, card *model.BankCard) (*model.BankCard, error)
	// GetBankCard - return bank card for user
	GetBankCard(ctx context.Context, userID string, cardID string) (*model.BankCard, error)
	// GetBankCardList - return all bank cards for user
	GetBankCardList(ctx context.Context, userID string) ([]model.BankCard, error)
	// AddBinary - adds new binaey data for user
	AddBinary(ctx context.Context, data *model.Binary) (*model.Binary, error)
	// GetBinary - return binary data for user
	GetBinary(ctx context.Context, userID string, binID string) (*model.Binary, error)
	// GetBinaryList - return all binary data for user
	GetBinaryList(ctx context.Context, userID string) ([]model.Binary, error)
}

// HttpHandler - handler for http router
type HttpHandler struct {
	auth    *auth.Auth
	usecase Usecase
}

// NewHttpHandler - return new instance
func NewHttpHandler(usecase Usecase) *HttpHandler {
	return &HttpHandler{
		auth:    auth.NewAuth(),
		usecase: usecase,
	}
}
