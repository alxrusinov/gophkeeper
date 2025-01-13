package notehandler

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// NoteHandler - bankcard handler
type NoteHandler struct {
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

	// GetNote - return note for user by note id
	GetNote(ctx context.Context, userID string, noteID string) (*model.Note, error)
	// GetNotes - return notes for user by note id
	GetNoteList(ctx context.Context, userID string) ([]model.Note, error)
	// AddNote - adds new note for user
	AddNote(ctx context.Context, note *model.Note) (*model.Note, error)

	// DeleteBankCard - delete note
	DeleteNote(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
}

// NewNoteHandlerHandler - crete new instance of bank card structure
func NewNoteHandlerHandler(usecase Usecase, auth Auth) *NoteHandler {
	return &NoteHandler{
		auth:    auth,
		usecase: usecase,
	}
}
