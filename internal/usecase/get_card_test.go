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

func TestUsecase_GetBankCard(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	successUserID := primitive.NewObjectID().Hex()
	successCardID := primitive.NewObjectID().Hex()

	successResult := &model.BankCard{ID: primitive.NewObjectID().Hex(), UserID: successUserID, Title: "Title", Data: 123, Meta: "meta"}

	repo.On("GetBankCard", mock.Anything, mock.Anything, mock.Anything).Return(successResult, nil)

	type args struct {
		ctx    context.Context
		userID string
		cardID string
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    *model.BankCard
		wantErr bool
	}{
		{
			name: "1# success",
			u:    testusecase,
			args: args{
				ctx:    context.Background(),
				userID: successUserID,
				cardID: successCardID,
			},
			want:    successResult,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetBankCard(tt.args.ctx, tt.args.userID, tt.args.cardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
