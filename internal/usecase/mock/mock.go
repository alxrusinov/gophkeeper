package mock

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/stretchr/testify/mock"
)

const (
	indexZero int = iota
	indexFirst
	indexSecond
)

// UsecaseMock - mock of usecase
type UsecaseMock struct {
	mock.Mock
}

// CreateUser - create new user
func (uc *UsecaseMock) CreateUser(ctx context.Context, lg *model.Login) (*model.User, error) {
	args := uc.Called(ctx, lg)

	return args.Get(indexZero).(*model.User), args.Error(indexFirst)
}

// VerifyUser - checks if user exists and has valid password
func (uc *UsecaseMock) VerifyUser(ctx context.Context, lg *model.Login) (*model.User, error) {
	args := uc.Called(ctx, lg)

	return args.Get(indexZero).(*model.User), args.Error(indexFirst)
}

// AddNote - adds new note for user
func (uc *UsecaseMock) AddNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	args := uc.Called(ctx, note)

	return args.Get(indexZero).(*model.Note), args.Error(indexFirst)
}

// GetNote - returns note for user by note id
func (uc *UsecaseMock) GetNote(ctx context.Context, userID, noteID string) (*model.Note, error) {
	args := uc.Called(ctx, userID, noteID)

	return args.Get(indexZero).(*model.Note), args.Error(indexFirst)
}

// GetNoteList - return notes for user by note id
func (uc *UsecaseMock) GetNoteList(ctx context.Context, userID string) ([]model.Note, error) {
	args := uc.Called(ctx, userID)

	return args.Get(indexZero).([]model.Note), args.Error(indexFirst)
}

// AddCredentials - adds new credentials for user
func (uc *UsecaseMock) AddCredentials(ctx context.Context, creds *model.Credentials) (*model.Credentials, error) {
	args := uc.Called(ctx, creds)

	return args.Get(indexZero).(*model.Credentials), args.Error(indexFirst)
}

// GetCredentilas - return credentials for user
func (uc *UsecaseMock) GetCredentials(ctx context.Context, userID string, credsID string) (*model.Credentials, error) {
	args := uc.Called(ctx, userID, credsID)

	return args.Get(indexZero).(*model.Credentials), args.Error(indexFirst)
}

// GetCredentialsList - return all credentials for user
func (uc *UsecaseMock) GetCredentialsList(ctx context.Context, userID string) ([]model.Credentials, error) {
	args := uc.Called(ctx, userID)

	return args.Get(indexZero).([]model.Credentials), args.Error(indexFirst)
}

// AddBankCard - adds new bank card for user
func (uc *UsecaseMock) AddBankCard(ctx context.Context, card *model.BankCard) (*model.BankCard, error) {
	args := uc.Called(ctx, card)

	return args.Get(indexZero).(*model.BankCard), args.Error(indexFirst)
}

// GetBankCard - return bank card for user
func (uc *UsecaseMock) GetBankCard(ctx context.Context, userID string, cardID string) (*model.BankCard, error) {
	args := uc.Called(ctx, userID, cardID)

	return args.Get(indexZero).(*model.BankCard), args.Error(indexFirst)
}

// GetBankCardList - return all bank cards for user
func (uc *UsecaseMock) GetBankCardList(ctx context.Context, userID string) ([]model.BankCard, error) {
	args := uc.Called(ctx, userID)

	return args.Get(indexZero).([]model.BankCard), args.Error(indexFirst)
}

// AddBinary - adds new binaey data for user
func (uc *UsecaseMock) AddBinary(ctx context.Context, data *model.Binary) (*model.Binary, error) {
	args := uc.Called(ctx, data)

	return args.Get(indexZero).(*model.Binary), args.Error(indexFirst)
}

// GetBinary - return binary data for user
func (uc *UsecaseMock) GetBinary(ctx context.Context, userID string, binID string) (*model.Binary, error) {
	args := uc.Called(ctx, userID, binID)

	return args.Get(indexZero).(*model.Binary), args.Error(indexFirst)
}

// GetBinaryList - return all binary data for user
func (uc *UsecaseMock) GetBinaryList(ctx context.Context, userID string) ([]model.Binary, error) {
	args := uc.Called(ctx, userID)

	return args.Get(indexZero).([]model.Binary), args.Error(indexFirst)
}

// DeleteBankCard - delete bank card
func (uc *UsecaseMock) DeleteBankCard(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	args := uc.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)

}

// DeleteBankCard - delete binary
func (uc *UsecaseMock) DeleteBinary(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {

	args := uc.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)
}

// DeleteBankCard - delete note
func (uc *UsecaseMock) DeleteNote(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	args := uc.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)

}

// DeleteBankCard - delete credentials
func (uc *UsecaseMock) DeleteCredentials(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	args := uc.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)

}

// CheckUser - checks if user from token existss in repository
func (uc *UsecaseMock) CheckUser(ctx context.Context, userID string) (bool, error) {
	args := uc.Called(ctx, userID)

	return args.Bool(indexZero), args.Error(indexFirst)
}

// NewUsecaseMock - returns new usecase mock
func NewUsecaseMock() *UsecaseMock {
	return new(UsecaseMock)
}
