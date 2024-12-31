package httphandler

import (
	"errors"
	"fmt"
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

func TestHttpHandler_GetNote(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	testAuth := authMock.NewAuthMock()
	testHandler := NewHttpHandler(testUsecase, testAuth)

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "User",
	}

	successNoteID := primitive.NewObjectID().Hex()
	notFoundID := ""
	errID := primitive.NewObjectID().Hex()
	noNote := primitive.NewObjectID().Hex()

	testAuth.On("GetVerifier").Return(new(jwt.Verifier))
	testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)

	app := iris.New()

	app.Get("/api/notes/{id}", testHandler.GetNote)

	app.Use(testHandler.AuthMiddleware())

	server := httptest.New(t, app, httptest.URL("http://example.com"), httptest.Debug(true), httptest.LogLevel("debug"))

	errAuth := authMock.NewAuthMock()

	errHandler := NewHttpHandler(testUsecase, errAuth)

	errApp := iris.New()

	errApp.Get("/api/notes/{id}", errHandler.GetNote)

	errServer := httptest.New(t, app, httptest.URL("http://err-example.com"))

	tests := []struct {
		name    string
		resCode int
		arg     string
		errAuth bool
	}{
		{
			name:    "1# success",
			resCode: http.StatusOK,
			arg:     successNoteID,
			errAuth: false,
		},
		{
			name:    "2# not found",
			resCode: http.StatusNotFound,
			arg:     notFoundID,
			errAuth: false,
		},
		{
			name:    "3# unauthorized",
			resCode: http.StatusUnauthorized,
			arg:     successNoteID,
			errAuth: true,
		},
		{
			name:    "4# error note",
			resCode: http.StatusInternalServerError,
			arg:     errID,
			errAuth: true,
		},
		{
			name:    "5# not note",
			resCode: http.StatusNotFound,
			arg:     noNote,
			errAuth: true,
		},
	}

	testUsecase.On("GetNote", mock.Anything, mock.Anything, successNoteID).Return(&model.Note{
		ID:     successNoteID,
		UserID: successUser.ID,
		Title:  "Title",
		Data:   "Note",
		Meta:   "Meta info",
	}, nil)
	testUsecase.On("GetNote", mock.Anything, mock.Anything, errID).Return(new(model.Note), errors.New("err"))
	testUsecase.On("GetNote", mock.Anything, mock.Anything, noNote).Return(new(model.Note), &customerrors.NotFound{Err: errors.New("err")})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case tests[0].name:
				server.GET(fmt.Sprintf("/api/notes/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[1].name:
				server.GET(fmt.Sprintf("/api/notes/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[2].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(new(model.User), errors.New("error of context"))
				errServer.GET(fmt.Sprintf("/api/notes/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[3].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)
				server.GET(fmt.Sprintf("/api/notes/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[4].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)
				server.GET(fmt.Sprintf("/api/notes/%s", tt.arg)).Expect().Status(tt.resCode)

			}

		})
	}
}
