package mock

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRepositoryMock_CreateUser(t *testing.T) {
	rp := NewRepositoryMock()

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
		rp      *RepositoryMock
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
			args: args{
				ctx: context.Background(),
				lg:  successLogin,
			},
			want:    successUser,
			wantErr: false,
		},
	}

	rp.On("CreateUser", mock.Anything, successLogin).Return(successUser, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rp.CreateUser(tt.args.ctx, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_AddNote(t *testing.T) {
	rp := NewRepositoryMock()
	successNote := &model.Note{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   "text",
		Meta:   "meta",
	}

	rp.On("AddNote", mock.Anything, successNote).Return(successNote, nil)

	type args struct {
		ctx  context.Context
		note *model.Note
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.Note
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.AddNote(tt.args.ctx, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.AddNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.AddNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_GetNoteList(t *testing.T) {
	rp := NewRepositoryMock()
	successNotes := []model.Note{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   "text",
		Meta:   "meta",
	}}

	successID := primitive.NewObjectID().Hex()

	rp.On("GetNoteList", mock.Anything, successID).Return(successNotes, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    []model.Note
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.GetNoteList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.GetNoteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.GetNoteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_AddCredentials(t *testing.T) {
	rp := NewRepositoryMock()
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

	rp.On("AddCredentials", mock.Anything, success).Return(success, nil)
	type args struct {
		ctx   context.Context
		creds *model.Credentials
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.Credentials
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.AddCredentials(tt.args.ctx, tt.args.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.AddCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.AddCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_GetCredentials(t *testing.T) {
	rp := NewRepositoryMock()

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
		rp      *RepositoryMock
		args    args
		want    *model.Credentials
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				credsID: successID,
			},
			want:    success,
			wantErr: false,
		},
	}
	rp.On("GetCredentials", mock.Anything, userID, successID).Return(success, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rp.GetCredentials(tt.args.ctx, tt.args.userID, tt.args.credsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.GetCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.GetCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_GetCredentialsList(t *testing.T) {
	rp := NewRepositoryMock()
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

	rp.On("GetCredentialsList", mock.Anything, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    []model.Credentials
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.GetCredentialsList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.GetCredentialsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.GetCredentialsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_AddBankCard(t *testing.T) {
	rp := NewRepositoryMock()
	success := &model.BankCard{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   123,
		Meta:   "meta",
	}

	rp.On("AddBankCard", mock.Anything, success).Return(success, nil)
	type args struct {
		ctx  context.Context
		card *model.BankCard
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.BankCard
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.AddBankCard(tt.args.ctx, tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.AddBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.AddBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_GetBankCard(t *testing.T) {
	rp := NewRepositoryMock()

	successID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	success := &model.BankCard{
		ID:     successID,
		UserID: userID,
		Title:  "Title",
		Data:   123,
		Meta:   "meta",
	}

	rp.On("GetBankCard", mock.Anything, userID, successID).Return(success, nil)

	type args struct {
		ctx    context.Context
		userID string
		cardID string
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.BankCard
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.GetBankCard(tt.args.ctx, tt.args.userID, tt.args.cardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.GetBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.GetBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_GetBankCardList(t *testing.T) {
	rp := NewRepositoryMock()
	success := []model.BankCard{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		Data:   123,
		Meta:   "meta",
	}}

	successID := primitive.NewObjectID().Hex()

	rp.On("GetBankCardList", mock.Anything, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    []model.BankCard
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.GetBankCardList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.GetBankCardList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.GetBankCardList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_AddBinary(t *testing.T) {
	// rp := NewRepositoryMock()
	// success := &model.Binary{
	// 	ID:     primitive.NewObjectID().Hex(),
	// 	UserID: primitive.NewObjectID().Hex(),
	// 	Title:  "Title",
	// 	Data:   []byte("foo"),
	// 	Meta:   "meta",
	// }

	// rp.On("AddBinary", mock.Anything, success).Return(success, nil)
	// type args struct {
	// 	ctx  context.Context
	// 	data *model.Binary
	// }
	// tests := []struct {
	// 	name    string
	// 	rp      *RepositoryMock
	// 	args    args
	// 	want    *model.Binary
	// 	wantErr bool
	// }{
	// 	{
	// 		name: "1# success",
	// 		rp:   rp,
	// 		args: args{
	// 			ctx:  context.Background(),
	// 			data: success,
	// 		},
	// 		want:    success,
	// 		wantErr: false,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := tt.rp.AddBinary(tt.args.ctx, tt.args.data)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("RepositoryMock.AddBinary() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("RepositoryMock.AddBinary() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestRepositoryMock_GetBinary(t *testing.T) {
	rp := NewRepositoryMock()

	successID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	success := &model.Binary{
		ID:     successID,
		UserID: userID,
		Title:  "Title",
		FileID: primitive.NewObjectID().Hex(),
		Meta:   "meta",
	}

	rp.On("GetBinary", mock.Anything, userID, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
		binID  string
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.Binary
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.GetBinary(tt.args.ctx, tt.args.userID, tt.args.binID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.GetBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.GetBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_GetBinaryList(t *testing.T) {
	rp := NewRepositoryMock()
	success := []model.Binary{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Title:  "Title",
		FileID: primitive.NewObjectID().Hex(),
		Meta:   "meta",
	}}

	successID := primitive.NewObjectID().Hex()

	rp.On("GetBinaryList", mock.Anything, successID).Return(success, nil)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    []model.Binary
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.GetBinaryList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.GetBinaryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.GetBinaryList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_DeleteBankCard(t *testing.T) {
	rp := NewRepositoryMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	rp.On("DeleteBankCard", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.DeleteBankCard(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.DeleteBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.DeleteBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRepositoryMock_DeleteCredentials(t *testing.T) {
	rp := NewRepositoryMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	rp.On("DeleteCredentials", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.DeleteCredentials(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.DeleteCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.DeleteCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRepositoryMock_DeleteBinary(t *testing.T) {
	rp := NewRepositoryMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	rp.On("DeleteBinary", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.DeleteBinary(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.DeleteBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.DeleteBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRepositoryMock_DeleteNote(t *testing.T) {
	rp := NewRepositoryMock()
	source := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	rp.On("DeleteNote", mock.Anything, source).Return(source, nil)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		rp      *RepositoryMock
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			rp:   rp,
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
			got, err := tt.rp.DeleteNote(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.DeleteNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMock.DeleteNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMock_CheckUser(t *testing.T) {
	rp := NewRepositoryMock()

	successUserID := primitive.NewObjectID().Hex()
	errUserID := primitive.NewObjectID().Hex()
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		rm      *RepositoryMock
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1# success",
			rm:   rp,
			args: args{
				ctx:    context.Background(),
				userID: successUserID,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "2# error",
			rm:   rp,
			args: args{
				ctx:    context.Background(),
				userID: errUserID,
			},
			want:    false,
			wantErr: true,
		},
	}

	rp.On("CheckUser", mock.Anything, successUserID).Return(true, nil)
	rp.On("CheckUser", mock.Anything, errUserID).Return(false, errors.New("err"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rm.CheckUser(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMock.CheckUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RepositoryMock.CheckUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
