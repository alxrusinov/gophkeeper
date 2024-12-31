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

func TestUsecase_AddBinary(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	resData := &model.Binary{
		ID: primitive.NewObjectID().Hex(),
	}
	errData := &model.Binary{
		ID: primitive.NewObjectID().Hex(),
	}

	err := errors.New("err")

	repo.On("AddBinary", mock.Anything, resData).Return(resData, nil)
	repo.On("AddBinary", mock.Anything, errData).Return(errData, err)

	type args struct {
		ctx  context.Context
		data *model.Binary
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
				ctx:  context.Background(),
				data: resData,
			},
			want:    resData,
			wantErr: false,
		},
		{
			name: "2# error",
			u:    testusecase,
			args: args{
				ctx:  context.Background(),
				data: errData,
			},
			want:    errData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.AddBinary(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.AddBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
