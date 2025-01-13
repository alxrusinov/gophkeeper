package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/model"
	mockRepo "github.com/alxrusinov/gophkeeper/internal/repository/mock"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUsecase_AddBankCard(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	resData := &model.BankCard{
		ID: primitive.NewObjectID().Hex(),
	}
	errData := &model.BankCard{
		ID: primitive.NewObjectID().Hex(),
	}

	err := errors.New("err")

	repo.On("AddBankCard", mock.Anything, resData).Return(resData, nil)
	repo.On("AddBankCard", mock.Anything, errData).Return(errData, err)
	type args struct {
		ctx  context.Context
		card *model.BankCard
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
				ctx:  context.Background(),
				card: resData,
			},
			want:    resData,
			wantErr: false,
		},
		{
			name: "2# error",
			u:    testusecase,
			args: args{
				ctx:  context.Background(),
				card: errData,
			},
			want:    errData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.AddBankCard(tt.args.ctx, tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.AddBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.AddBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
