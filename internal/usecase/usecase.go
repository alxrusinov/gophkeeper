package usecase

import "github.com/alxrusinov/gophkeeper/internal/model"

// Repository - interface of repository
type Repository interface {
	// CreateUser - create new user
	CreateUser(lg *model.Login) (*model.User, error)
	// VerifyUser - checks if user exists and has valid password
	VerifyUser(lg *model.Login) (*model.User, error)
	// AddNote - adds new note for user
	AddNote(note *model.Note, userID string) (*model.Note, error)
	// GetNote - returns note for user by note id
	GetNote(userID, noteID string) (*model.Note, error)
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
