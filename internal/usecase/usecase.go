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
	AddNote(ctx context.Context, note *model.Note, userID string) (*model.Note, error)
	// GetNote - returns note for user by note id
	GetNote(ctx context.Context, userID, noteID string) (*model.Note, error)
	// GetNotes - return notes for user by note id
	GetNoteList(ctx context.Context, userID string) ([]model.Note, error)
	// AddCredentials - adds new credentials for user
	AddCredentials(ctx context.Context, creds *model.Credentials, userID string) (*model.Credentials, error)
	// GetCredentilas - return credentials for user
	GetCredentials(ctx context.Context, userID string, credsID string) (*model.Credentials, error)
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
