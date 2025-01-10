package httphandler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/auth"
	authMock "github.com/alxrusinov/gophkeeper/internal/auth/mock"
	"github.com/alxrusinov/gophkeeper/internal/config"
	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHttpHandler_Login(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	cfg := config.NewConfig()
	cfg.RunMock()
	testAuth := auth.NewAuth(*cfg)
	testHandler := NewHttpHandler(testUsecase, testAuth)

	app := iris.New()

	app.Post("/auth/login", testHandler.Login)

	server := httptest.New(t, app, httptest.URL("http://example.com"))

	successLogin := &model.Login{
		Username: "success",
		Password: "111",
	}

	nonExistsLogin := &model.Login{
		Username: "nonexist",
		Password: "222",
	}

	undverifiedLogin := &model.Login{
		Username: "unverified",
		Password: "333",
	}

	badLogin := &model.Login{
		Username: "bad",
		Password: "444",
	}

	successUser := &model.User{
		Username: successLogin.Username,
		ID:       primitive.NewObjectID().Hex(),
	}

	testUsecase.On("VerifyUser", mock.Anything, successLogin).Return(successUser, nil)
	testUsecase.On("VerifyUser", mock.Anything, nonExistsLogin).Return(nil, customerrors.NonexistentUser{Err: errors.New("nonexist user")})
	testUsecase.On("VerifyUser", mock.Anything, undverifiedLogin).Return(nil, customerrors.UnverifiedUser{Err: errors.New("unverifed")})
	testUsecase.On("VerifyUser", mock.Anything, badLogin).Return(new(model.User), errors.New("error"))

	errAuth := authMock.NewAuthMock()

	errHandler := NewHttpHandler(testUsecase, errAuth)

	errApp := iris.New()

	errApp.Post("/auth/login", errHandler.Login)

	errServer := httptest.New(t, app, httptest.URL("http://err-example.com"))

	errAuth.On("GetTokenPair", mock.Anything).Return(nil, errors.New("err"))

	tests := []struct {
		name     string
		resCode  int
		arg      *model.Login
		tokenErr bool
		badReq   bool
	}{
		{
			name:     "1# success",
			resCode:  http.StatusOK,
			arg:      successLogin,
			tokenErr: false,
		},
		{
			name:     "2# bad request",
			resCode:  http.StatusBadRequest,
			arg:      successLogin,
			tokenErr: false,
			badReq:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch {
			case tt.tokenErr:
				errServer.POST("/auth/login").WithJSON(tt.arg).Expect().Status(tt.resCode)
			case tt.badReq:
				server.POST("/auth/login").WithJSON("clown").Expect().Status(tt.resCode)
			default:
				server.POST("/auth/login").WithJSON(tt.arg).Expect().Status(tt.resCode)

			}

		})
	}
}
