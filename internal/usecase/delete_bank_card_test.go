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

func TestUsecase_DeleteBankCard(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	resData := &model.SourceID{
		ID: primitive.NewObjectID().Hex(),
	}
	errData := &model.SourceID{
		ID: primitive.NewObjectID().Hex(),
	}

	err := errors.New("err")

	repo.On("DeleteBankCard", mock.Anything, resData).Return(resData, nil)
	repo.On("DeleteBankCard", mock.Anything, errData).Return(errData, err)
	type args struct {
		ctx    context.Context
		source *model.SourceID
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    *model.SourceID
		wantErr bool
	}{
		{
			name: "1# success",
			u:    testusecase,
			args: args{
				ctx:    context.Background(),
				source: resData,
			},
			want:    resData,
			wantErr: false,
		},
		{
			name: "2# error",
			u:    testusecase,
			args: args{
				ctx:    context.Background(),
				source: errData,
			},
			want:    errData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.DeleteBankCard(tt.args.ctx, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.DeleteBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.DeleteBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
