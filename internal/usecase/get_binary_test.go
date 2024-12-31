package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/model"
	mockRepo "github.com/alxrusinov/gophkeeper/internal/repository/mock"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUsecase_GetBinary(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	successUserID := primitive.NewObjectID().Hex()
	successBinID := primitive.NewObjectID().Hex()

	successResult := &model.Binary{ID: primitive.NewObjectID().Hex(), UserID: successUserID, Data: []byte("string"), Meta: "meta"}

	repo.On("GetBinaryList", mock.Anything, mock.Anything).Return(successResult, nil)

	type args struct {
		ctx    context.Context
		userID string
		binID  string
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    *model.Binary
		wantErr bool
	}{
		{
			name: "1# success",
			u:    testusecase,
			args: args{
				ctx:    context.Background(),
				userID: successUserID,
				binID:  successBinID,
			},
			want:    successResult,
			wantErr: false,
		},
	}

	repo.On("GetBinary", mock.Anything, mock.Anything, mock.Anything).Return(successResult, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetBinary(tt.args.ctx, tt.args.userID, tt.args.binID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
