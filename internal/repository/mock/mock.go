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

// RepositoryMock - mock of repository
type RepositoryMock struct {
	mock.Mock
}

// CreateUser - create new user
func (rm *RepositoryMock) CreateUser(ctx context.Context, lg *model.Login) (*model.User, error) {
	args := rm.Called(ctx, lg)

	return args.Get(indexZero).(*model.User), args.Error(indexFirst)
}

// VerifyUser - checks if user exists and has valid password
func (rm *RepositoryMock) VerifyUser(ctx context.Context, lg *model.Login) (*model.User, error) {
	args := rm.Called(ctx, lg)

	return args.Get(indexZero).(*model.User), args.Error(indexFirst)
}

// AddNote - adds new note for user
func (rm *RepositoryMock) AddNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	args := rm.Called(ctx, note)

	return args.Get(indexZero).(*model.Note), args.Error(indexFirst)
}

// GetNote - returns note for user by note id
func (rm *RepositoryMock) GetNote(ctx context.Context, userID, noteID string) (*model.Note, error) {
	args := rm.Called(ctx, userID, noteID)

	return args.Get(indexZero).(*model.Note), args.Error(indexFirst)
}

// GetNotes - return notes for user by note id
func (rm *RepositoryMock) GetNoteList(ctx context.Context, userID string) ([]model.Note, error) {
	args := rm.Called(ctx, userID)

	return args.Get(indexZero).([]model.Note), args.Error(indexFirst)
}

// AddCredentials - adds new credentials for user
func (rm *RepositoryMock) AddCredentials(ctx context.Context, creds *model.Credentials) (*model.Credentials, error) {
	args := rm.Called(ctx, creds)

	return args.Get(indexZero).(*model.Credentials), args.Error(indexFirst)
}

// GetCredentilas - return credentials for user
func (rm *RepositoryMock) GetCredentials(ctx context.Context, userID string, credsID string) (*model.Credentials, error) {
	args := rm.Called(ctx, userID, credsID)

	return args.Get(indexZero).(*model.Credentials), args.Error(indexFirst)
}

// GetCredentialsList - return all credentials for user
func (rm *RepositoryMock) GetCredentialsList(ctx context.Context, userID string) ([]model.Credentials, error) {
	args := rm.Called(ctx, userID)

	return args.Get(indexZero).([]model.Credentials), args.Error(indexFirst)
}

// AddBankCard - adds new bank card for user
func (rm *RepositoryMock) AddBankCard(ctx context.Context, card *model.BankCard) (*model.BankCard, error) {
	args := rm.Called(ctx, card)

	return args.Get(indexZero).(*model.BankCard), args.Error(indexFirst)
}

// GetBankCard - return bank card for user
func (rm *RepositoryMock) GetBankCard(ctx context.Context, userID string, cardID string) (*model.BankCard, error) {
	args := rm.Called(ctx, userID, cardID)

	return args.Get(indexZero).(*model.BankCard), args.Error(indexFirst)
}

// GetBankCardList - return all bank cards for user
func (rm *RepositoryMock) GetBankCardList(ctx context.Context, userID string) ([]model.BankCard, error) {
	args := rm.Called(ctx, userID)

	return args.Get(indexZero).([]model.BankCard), args.Error(indexFirst)
}

// AddBinary - adds new binaey data for user
func (rm *RepositoryMock) AddBinary(ctx context.Context, data *model.Binary) (*model.Binary, error) {
	args := rm.Called(ctx, data)

	return args.Get(indexZero).(*model.Binary), args.Error(indexFirst)
}

// GetBinary - return binary data for user
func (rm *RepositoryMock) GetBinary(ctx context.Context, userID string, binID string) (*model.Binary, error) {
	args := rm.Called(ctx, userID, binID)

	return args.Get(indexZero).(*model.Binary), args.Error(indexFirst)
}

// GetBinaryList - return all binary data for user
func (rm *RepositoryMock) GetBinaryList(ctx context.Context, userID string) ([]model.Binary, error) {
	args := rm.Called(ctx, userID)

	return args.Get(indexZero).([]model.Binary), args.Error(indexFirst)
}

// DeleteBankCard - delete bank card
func (rm *RepositoryMock) DeleteBankCard(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	args := rm.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)

}

// DeleteBankCard - delete binary
func (rm *RepositoryMock) DeleteBinary(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {

	args := rm.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)
}

// DeleteBankCard - delete note
func (rm *RepositoryMock) DeleteNote(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	args := rm.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)

}

// DeleteBankCard - delete credentials
func (rm *RepositoryMock) DeleteCredentials(ctx context.Context, source *model.SourceID) (*model.SourceID, error) {
	args := rm.Called(ctx, source)

	return args.Get(indexZero).(*model.SourceID), args.Error(indexFirst)

}

// CheckUser - checks if user from token existss in repository
func (rm *RepositoryMock) CheckUser(ctx context.Context, userID string) (bool, error) {
	args := rm.Called(ctx, userID)

	return args.Bool(indexZero), args.Error(indexFirst)
}

// NewRepositoryMock - returns new repository mock
func NewRepositoryMock() *RepositoryMock {
	return new(RepositoryMock)
}
