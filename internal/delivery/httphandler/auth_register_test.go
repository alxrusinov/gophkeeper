package httphandler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/auth"
	authMock "github.com/alxrusinov/gophkeeper/internal/auth/mock"
	"github.com/alxrusinov/gophkeeper/internal/config"
	"github.com/alxrusinov/gophkeeper/internal/model"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHttpHandler_Register(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	cfg := config.NewConfig()
	cfg.RunMock()
	testAuth := auth.NewAuth(*cfg)
	testHandler := NewHttpHandler(testUsecase, testAuth)

	app := iris.New()

	app.Post("/auth/register", testHandler.Register)

	server := httptest.New(t, app, httptest.URL("http://example.com"))

	testUsecase.On("CreateUser", mock.Anything, mock.Anything).Return(&model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: primitive.NewObjectID().Hex(),
	}, nil)

	errAuth := authMock.NewAuthMock()

	errAuth.On("GetAccessToken", mock.Anything).Return("", errors.New("token error"))

	errHandler := NewHttpHandler(testUsecase, errAuth)

	errApp := iris.New()

	errApp.Post("/auth/register", errHandler.Register)

	errServer := httptest.New(t, errApp, httptest.URL("http://err-example.com"))

	tests := []struct {
		name    string
		resCode int
		arg     model.Login
	}{
		{
			name:    "1# success",
			resCode: http.StatusCreated,
			arg: model.Login{
				Username: "vasya",
				Password: "1111",
			},
		},
		{
			name:    "2# bad request",
			resCode: http.StatusBadRequest,
			arg: model.Login{
				Username: "vasya",
				Password: "1111",
			},
		},
		{
			name:    "3# error creating user",
			resCode: http.StatusInternalServerError,
			arg: model.Login{
				Username: "vasya",
				Password: "1111",
			},
		},
		{
			name:    "4# error token",
			resCode: http.StatusInternalServerError,
			arg: model.Login{
				Username: "vasya",
				Password: "1111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case tests[0].name:
				server.POST("/auth/register").WithJSON(tt.arg).Expect().Status(tt.resCode)
			case tests[1].name:
				server.POST("/auth/register").WithJSON("clown").Expect().Status(tt.resCode)
			case tests[2].name:
				testUsecase.On("CreateUser", mock.Anything, mock.Anything).Unset()
				testUsecase.On("CreateUser", mock.Anything, mock.Anything).Return(&model.User{
					ID:       primitive.NewObjectID().Hex(),
					Username: primitive.NewObjectID().Hex(),
				}, errors.New("some errors"))
				server.POST("/auth/register").WithJSON(tt.arg).Expect().Status(tt.resCode)
			case tests[3].name:
				testUsecase.On("CreateUser", mock.Anything, mock.Anything).Unset()
				testUsecase.On("CreateUser", mock.Anything, mock.Anything).Return(&model.User{
					ID:       primitive.NewObjectID().Hex(),
					Username: primitive.NewObjectID().Hex(),
				}, nil)
				errAuth.On("GetAccessToken", mock.Anything).Return("", errors.New("token failed"))
				errServer.POST("/auth/register").WithJSON(tt.arg).Expect().Status(tt.resCode)
			}
		})
	}
}
