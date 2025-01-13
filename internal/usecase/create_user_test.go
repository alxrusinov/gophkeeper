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

func TestUsecase_CreateUser(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	reqData := &model.Login{
		Username: primitive.NewObjectID().Hex(),
	}
	reqErrData := &model.Login{
		Username: primitive.NewObjectID().Hex(),
	}
	resData := &model.User{
		ID: primitive.NewObjectID().Hex(),
	}
	errData := &model.User{
		ID: primitive.NewObjectID().Hex(),
	}

	err := errors.New("err")

	repo.On("CreateUser", mock.Anything, reqData).Return(resData, nil)
	repo.On("CreateUser", mock.Anything, reqErrData).Return(errData, err)
	type args struct {
		ctx context.Context
		lg  *model.Login
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "1# success",
			u:    testusecase,
			args: args{
				ctx: context.Background(),
				lg:  reqData,
			},
			want:    resData,
			wantErr: false,
		},
		{
			name: "2# error",
			u:    testusecase,
			args: args{
				ctx: context.Background(),
				lg:  reqErrData,
			},
			want:    errData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.CreateUser(tt.args.ctx, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
