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
	GetNotes(ctx context.Context, userID string) ([]model.Note, error)
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
