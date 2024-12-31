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

func TestUsecase_GetCredentialsList(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	successUserID := primitive.NewObjectID().Hex()

	successResult := []model.Credentials{
		{ID: primitive.NewObjectID().Hex(), UserID: successUserID, Title: "Title", Data: model.Login{
			Username: "name",
			Password: "13",
		}, Meta: "meta"},
	}

	repo.On("GetCredentialsList", mock.Anything, mock.Anything).Return(successResult, nil)

	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    []model.Credentials
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
			got, err := tt.u.GetCredentialsList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetCredentialsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetCredentialsList() = %v, want %v", got, tt.want)
			}
		})
	}
}
