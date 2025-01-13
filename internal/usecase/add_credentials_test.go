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

func TestUsecase_AddCredentials(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	resData := &model.Credentials{
		ID: primitive.NewObjectID().Hex(),
	}
	errData := &model.Credentials{
		ID: primitive.NewObjectID().Hex(),
	}

	err := errors.New("err")

	repo.On("AddCredentials", mock.Anything, resData).Return(resData, nil)
	repo.On("AddCredentials", mock.Anything, errData).Return(errData, err)
	type args struct {
		ctx   context.Context
		creds *model.Credentials
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    *model.Credentials
		wantErr bool
	}{
		{
			name: "1# success",
			u:    testusecase,
			args: args{
				ctx:   context.Background(),
				creds: resData,
			},
			want:    resData,
			wantErr: false,
		},
		{
			name: "2# error",
			u:    testusecase,
			args: args{
				ctx:   context.Background(),
				creds: errData,
			},
			want:    errData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.AddCredentials(tt.args.ctx, tt.args.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.AddCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.AddCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
