package mock

import (
	"context"
	"reflect"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUsecaseMock_CreateUser(t *testing.T) {
	uc := NewUsecaseMock()

	successLogin := &model.Login{
		Username: "success",
		Password: "success",
	}

	successUser := &model.User{
		Username: successLogin.Username,
		ID:       primitive.NewObjectID().Hex(),
	}

	type args struct {
		ctx context.Context
		lg  *model.Login
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx: context.Background(),
				lg:  successLogin,
			},
			want:    successUser,
			wantErr: false,
		},
	}

	uc.On("CreateUser", mock.Anything, successLogin).Return(successUser, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CreateUser(tt.args.ctx, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_AddNote(t *testing.T) {
	uc := NewUsecaseMock()
	successNote := &model.Note{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   "text",
		Meta:   "meta",
	}

	uc.On("AddNote", mock.Anything, successNote).Return(successNote, nil)

	type args struct {
		ctx  context.Context
		note *model.Note
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.Note
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:  context.Background(),
				note: successNote,
			},
			want:    successNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.AddNote(tt.args.ctx, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.AddNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.AddNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_GetNoteList(t *testing.T) {
	uc := NewUsecaseMock()
	successNotes := []model.Note{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   "text",
		Meta:   "meta",
	}}

	successID := primitive.NewObjectID().Hex()

	uc.On("GetNoteList", mock.Anything, successID).Return(successNotes, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    []model.Note
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				userID: successID,
			},
			want:    successNotes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetNoteList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.GetNoteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.GetNoteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_AddCredentials(t *testing.T) {
	uc := NewUsecaseMock()
	success := &model.Credentials{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data: model.Login{
			Username: "name",
			Password: "passw",
		},
		Meta: "meta",
	}

	uc.On("AddCredentials", mock.Anything, success).Return(success, nil)
	type args struct {
		ctx   context.Context
		creds *model.Credentials
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.Credentials
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:   context.Background(),
				creds: success,
			},
			want:    success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.AddCredentials(tt.args.ctx, tt.args.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.AddCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.AddCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_GetCredentials(t *testing.T) {
	uc := NewUsecaseMock()

	successID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	success := &model.Credentials{
		ID:     successID,
		UserID: userID,
		Title:  "Title",
		Data: model.Login{
			Username: "name",
			Password: "passw",
		},
		Meta: "meta",
	}
	type args struct {
		ctx     context.Context
		userID  string
		credsID string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.Credentials
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				credsID: successID,
			},
			want:    success,
			wantErr: false,
		},
	}
	uc.On("GetCredentials", mock.Anything, userID, successID).Return(success, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetCredentials(tt.args.ctx, tt.args.userID, tt.args.credsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.GetCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.GetCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_GetCredentialsList(t *testing.T) {
	uc := NewUsecaseMock()
	success := []model.Credentials{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data: model.Login{
			Username: "name",
			Password: "passw",
		},
		Meta: "meta",
	}}

	successID := primitive.NewObjectID().Hex()

	uc.On("GetCredentialsList", mock.Anything, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    []model.Credentials
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				userID: successID,
			},
			want:    success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetCredentialsList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.GetCredentialsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.GetCredentialsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_AddBankCard(t *testing.T) {
	uc := NewUsecaseMock()
	success := &model.BankCard{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   123,
		Meta:   "meta",
	}

	uc.On("AddBankCard", mock.Anything, success).Return(success, nil)
	type args struct {
		ctx  context.Context
		card *model.BankCard
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.BankCard
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:  context.Background(),
				card: success,
			},
			want:    success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.AddBankCard(tt.args.ctx, tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.AddBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.AddBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_GetBankCard(t *testing.T) {
	uc := NewUsecaseMock()

	successID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	success := &model.BankCard{
		ID:     successID,
		UserID: userID,
		Title:  "Title",
		Data:   123,
		Meta:   "meta",
	}

	uc.On("GetBankCard", mock.Anything, userID, successID).Return(success, nil)

	type args struct {
		ctx    context.Context
		userID string
		cardID string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.BankCard
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				userID: userID,
				cardID: successID,
			},
			want:    success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetBankCard(tt.args.ctx, tt.args.userID, tt.args.cardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.GetBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.GetBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_GetBankCardList(t *testing.T) {
	uc := NewUsecaseMock()
	success := []model.BankCard{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   123,
		Meta:   "meta",
	}}

	successID := primitive.NewObjectID().Hex()

	uc.On("GetBankCardList", mock.Anything, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    []model.BankCard
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				userID: successID,
			},
			want:    success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetBankCardList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.GetBankCardList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.GetBankCardList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_AddBinary(t *testing.T) {
	uc := NewUsecaseMock()
	successWant := &model.Binary{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		FileID: primitive.NewObjectID().Hex(),
		Meta:   "meta",
	}
	success := &model.BinaryUpload{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		FileID: primitive.NewObjectID().Hex(),
		Meta:   "meta",
	}

	uc.On("AddBinary", mock.Anything, success).Return(successWant, nil)
	type args struct {
		ctx  context.Context
		data *model.Binary
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.Binary
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:  context.Background(),
				data: successWant,
			},
			want:    successWant,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.AddBinary(tt.args.ctx, success)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.AddBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_GetBinary(t *testing.T) {
	uc := NewUsecaseMock()

	successID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	success := &model.Binary{
		ID:     successID,
		UserID: userID,
		Title:  "Title",
		FileID: primitive.NewObjectID().Hex(),
		Meta:   "meta",
	}

	uc.On("GetBinary", mock.Anything, userID, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
		binID  string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.Binary
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				userID: userID,
				binID:  successID,
			},
			want:    success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetBinary(tt.args.ctx, tt.args.userID, tt.args.binID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.GetBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.GetBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_GetBinaryList(t *testing.T) {
	uc := NewUsecaseMock()
	success := []model.Binary{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		FileID: primitive.NewObjectID().Hex(),
		Meta:   "meta",
	}}

	successID := primitive.NewObjectID().Hex()

	uc.On("GetBinaryList", mock.Anything, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    []model.Binary
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				userID: successID,
			},
			want:    success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetBinaryList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.GetBinaryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.GetBinaryList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_DeleteBankCard(t *testing.T) {
	uc := NewUsecaseMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	uc.On("DeleteBankCard", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				source: source,
			},
			want:    source,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DeleteBankCard(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.DeleteBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.DeleteBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUsecaseMock_DeleteCredentials(t *testing.T) {
	uc := NewUsecaseMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	uc.On("DeleteCredentials", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				source: source,
			},
			want:    source,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DeleteCredentials(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.DeleteCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.DeleteCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUsecaseMock_DeleteBinary(t *testing.T) {
	uc := NewUsecaseMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	uc.On("DeleteBinary", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				source: source,
			},
			want:    source,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DeleteBinary(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.DeleteBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.DeleteBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUsecaseMock_DeleteNote(t *testing.T) {
	uc := NewUsecaseMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	uc.On("DeleteNote", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			uc:   uc,
			args: args{
				ctx:    context.Background(),
				source: source,
			},
			want:    source,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DeleteNote(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.DeleteNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecaseMock.DeleteNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseMock_CheckUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		uc      *UsecaseMock
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CheckUser(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMock.CheckUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UsecaseMock.CheckUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
