package usecase

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
)

// Repository - interface of repository
type Repository interface {
	// CreateUser - create new user
	CreateUser(ctx context.Context, lg *model.Login) (*model.User, error)
	// VerifyUser - checks if user exists and has valid password
	VerifyUser(ctx context.Context, lg *model.Login) (*model.User, error)
	// AddNote - adds new note for user
	AddNote(ctx context.Context, note *model.Note) (*model.Note, error)
	// GetNote - returns note for user by note id
	GetNote(ctx context.Context, userID, noteID string) (*model.Note, error)
	// GetNotes - return notes for user by note id
	GetNoteList(ctx context.Context, userID string) ([]model.Note, error)
	// AddCredentials - adds new credentials for user
	AddCredentials(ctx context.Context, creds *model.Credentials) (*model.Credentials, error)
	// GetCredentilas - return credentials for user
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
	// DeleteBankCard - delete bank card
	DeleteBankCard(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
	// DeleteBankCard - delete binary
	DeleteBinary(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
	// DeleteBankCard - delete note
	DeleteNote(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
	// DeleteBankCard - delete credentials
	DeleteCredentials(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
	// CheckUser - checks if user from token existss in repository
	CheckUser(ctx context.Context, userID string) (bool, error)
}

// Usecase implements httphandler.Usecase interface
type Usecase struct {
	repository Repository
}

// NewUsecase - create new instance of Usecase
func NewUsecase(repository Repository) *Usecase {
	uc := &Usecase{repository: repository}

	return uc
}
