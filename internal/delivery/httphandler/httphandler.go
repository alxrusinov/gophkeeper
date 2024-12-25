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
	GetNotes(ctx context.Context, userID string) ([]model.Note, error)
	// AddNote - adds new note for user
	AddNote(ctx context.Context, note *model.Note, userID string) (*model.Note, error)
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