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

func TestUsecase_GetCardList(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	successUserID := primitive.NewObjectID().Hex()

	successResult := []model.BankCard{{ID: primitive.NewObjectID().Hex(), UserID: successUserID, Title: "Title", Data: 123, Meta: "meta"}}

	repo.On("GetBankCardList", mock.Anything, mock.Anything).Return(successResult, nil)

	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    []model.BankCard
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

	repo.On("GetBinary", mock.Anything, mock.Anything, mock.Anything).Return(successResult, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetBankCardList(tt.args.ctx, tt.args.userID)
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
