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

func TestUsecase_GetBinaryList(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)
	type args struct {
		ctx    context.Context
		userID string
	}

	successUserID := primitive.NewObjectID().Hex()

	successResult := []model.Binary{
		{ID: primitive.NewObjectID().Hex(), UserID: successUserID, FileID: primitive.NewObjectID().Hex()}}

	repo.On("GetBinaryList", mock.Anything, mock.Anything).Return(successResult, nil)

	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    []model.Binary
		wantErr bool
	}{
		{
			name: "1# success",
			u:    testusecase,
			args: args{
				ctx:    context.Background(),
				userID: successUserID,
			},
			want:    successResult,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetBinaryList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetBinaryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetBinaryList() = %v, want %v", got, tt.want)
			}
		})
	}
}
