package httphandler

import (
	"errors"
	"net/http"
	"testing"

	authMock "github.com/alxrusinov/gophkeeper/internal/auth/mock"
	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHttpHandler_GetNoteList(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	testAuth := authMock.NewAuthMock()
	testHandler := NewHttpHandler(testUsecase, testAuth)

	successSource := []model.Note{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Data:   "text",
		Title:  "title",
		Meta:   "meta",
	}}

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "User",
	}

	testUsecase.On("GetNoteList", mock.Anything, mock.Anything).Return(successSource, nil)

	testAuth.On("GetVerifier").Return(new(jwt.Verifier))
	testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)

	app := iris.New()

	app.Get("/api/notes", testHandler.GetNoteList)

	app.Use(testHandler.AuthMiddleware())

	server := httptest.New(t, app, httptest.URL("http://example.com"), httptest.Debug(true), httptest.LogLevel("debug"))

	tests := []struct {
		name    string
		resCode int
		source  []model.Note
		userID  string
	}{
		{
			name:    "1# success",
			resCode: http.StatusOK,
			source:  successSource,
			userID:  successUser.ID,
		},
		{
			name:    "2# not found",
			resCode: http.StatusNotFound,
			source:  successSource,
			userID:  successUser.ID,
		},
		{
			name:    "3# unauthorized error",
			resCode: http.StatusUnauthorized,
			source:  successSource,
			userID:  successUser.ID,
		},
		{
			name:    "4# error",
			resCode: http.StatusInternalServerError,
			source:  successSource,
			userID:  successUser.ID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case tests[0].name:
				server.GET("/api/notes").Expect().Status(tt.resCode)
			case tests[1].name:
				testUsecase.On("GetNoteList", mock.Anything, mock.Anything).Unset()
				testUsecase.On("GetNoteList", mock.Anything, mock.Anything).Return(tt.source, &customerrors.NotFound{})
				server.GET("/api/notes").Expect().Status(tt.resCode)
			case tests[2].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, errors.New("error"))
				server.GET("/api/notes").Expect().Status(tt.resCode)
			case tests[3].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)
				testUsecase.On("GetNoteList", mock.Anything, mock.Anything).Unset()
				testUsecase.On("GetNoteList", mock.Anything, mock.Anything).Return(tt.source, errors.New("error"))
				server.GET("/api/notes").Expect().Status(tt.resCode)

			}
		})
	}
}
