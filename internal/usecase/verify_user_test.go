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

func TestUsecase_VerifyUser(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	successLogin := model.Login{
		Username: "success",
		Password: "success",
	}

	errLogin := model.Login{
		Username: "Err",
		Password: "Err",
	}

	successUser := model.User{
		Username: successLogin.Username,
		ID:       primitive.NewObjectID().Hex(),
	}

	errUser := model.User{
		Username: errLogin.Username,
		ID:       primitive.NewObjectID().Hex(),
	}

	repo.On("VerifyUser", mock.Anything, &successLogin).Return(&successUser, nil)
	repo.On("VerifyUser", mock.Anything, &errLogin).Return(&errUser, errors.New("error"))

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
				lg:  &successLogin,
			},
			want:    &successUser,
			wantErr: false,
		},
		{
			name: "2# error",
			u:    testusecase,
			args: args{
				ctx: context.Background(),
				lg:  &errLogin,
			},
			want:    &errUser,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.VerifyUser(tt.args.ctx, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.VerifyUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.VerifyUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
