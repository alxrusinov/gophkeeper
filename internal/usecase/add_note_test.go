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

func TestUsecase_AddNote(t *testing.T) {
	repo := mockRepo.NewRepositoryMock()

	testusecase := NewUsecase(repo)

	resData := &model.Note{
		ID: primitive.NewObjectID().Hex(),
	}
	errData := &model.Note{
		ID: primitive.NewObjectID().Hex(),
	}

	err := errors.New("err")

	repo.On("AddNote", mock.Anything, resData).Return(resData, nil)
	repo.On("AddNote", mock.Anything, errData).Return(errData, err)
	type args struct {
		ctx  context.Context
		note *model.Note
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    *model.Note
		wantErr bool
	}{
		{
			name: "1# success",
			u:    testusecase,
			args: args{
				ctx:  context.Background(),
				note: resData,
			},
			want:    resData,
			wantErr: false,
		},
		{
			name: "2# error",
			u:    testusecase,
			args: args{
				ctx:  context.Background(),
				note: errData,
			},
			want:    errData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.AddNote(tt.args.ctx, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.AddNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.AddNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
