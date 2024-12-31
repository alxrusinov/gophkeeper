package httphandler

import (
	"errors"
	"net/http"
	"testing"

	authMock "github.com/alxrusinov/gophkeeper/internal/auth/mock"
	"github.com/alxrusinov/gophkeeper/internal/model"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHttpHandler_DeleteCredentials(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	testAuth := authMock.NewAuthMock()
	testHandler := NewHttpHandler(testUsecase, testAuth)

	successSource := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "User",
	}

	testUsecase.On("DeleteCredentials", mock.Anything, mock.Anything).Return(successSource, nil)

	testAuth.On("GetVerifier").Return(new(jwt.Verifier))
	testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)

	app := iris.New()

	app.Delete("/api/credentials", testHandler.DeleteCredentials)

	app.Use(testHandler.AuthMiddleware())

	server := httptest.New(t, app, httptest.URL("http://example.com"), httptest.Debug(true), httptest.LogLevel("debug"))

	tests := []struct {
		name    string
		resCode int
		source  *model.SourceID
	}{
		{
			name:    "1# success",
			resCode: http.StatusOK,
			source:  successSource,
		},
		{
			name:    "2# bad request",
			resCode: http.StatusBadRequest,
			source:  successSource,
		},
		{
			name:    "3# unauthorized error",
			resCode: http.StatusUnauthorized,
			source:  successSource,
		},
		{
			name:    "4# error",
			resCode: http.StatusInternalServerError,
			source:  successSource,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case tests[0].name:
				server.DELETE("/api/credentials").WithJSON(tt.source).Expect().Status(tt.resCode)
			case tests[1].name:
				server.DELETE("/api/credentials").WithJSON("foo").Expect().Status(tt.resCode)
			case tests[2].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, errors.New("error"))
				server.DELETE("/api/credentials").WithJSON(tt.source).Expect().Status(tt.resCode)
			case tests[3].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)
				testUsecase.On("DeleteCredentials", mock.Anything, mock.Anything).Unset()
				testUsecase.On("DeleteCredentials", mock.Anything, mock.Anything).Return(tt.source, errors.New("error"))
				server.DELETE("/api/credentials").WithJSON(tt.source).Expect().Status(tt.resCode)

			}
		})
	}
}
