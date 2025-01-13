package binaryhandler

import (
	"bytes"
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// BinaryHandler - binary handler
type BinaryHandler struct {
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
	// AddBinary - adds new binaey data for user
	AddBinary(ctx context.Context, data *model.BinaryUpload) (*model.Binary, error)
	// GetBinary - return binary data for user
	GetBinary(ctx context.Context, userID string, binID string) (*model.Binary, error)
	// GetBinaryList - return all binary data for user
	GetBinaryList(ctx context.Context, userID string) ([]model.Binary, error)

	// DeleteBankCard - delete binary
	DeleteBinary(ctx context.Context, source *model.SourceID) (*model.SourceID, error)
	// DownloadFile - downloads file by ud
	DownloadFile(ctx context.Context, fileID string) (*bytes.Buffer, error)
}

// NewBinaryHandlerHandler - crete new instance of binary structure
func NewBinaryHandlerHandler(usecase Usecase, auth Auth) *BinaryHandler {
	return &BinaryHandler{
		auth:    auth,
		usecase: usecase,
	}
}
