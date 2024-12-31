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

func TestHttpHandler_SetNote(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	testAuth := authMock.NewAuthMock()
	testHandler := NewHttpHandler(testUsecase, testAuth)

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "User",
	}

	source := &model.Note{
		ID:     primitive.NewObjectID().Hex(),
		UserID: successUser.ID,
		Title:  "title",
		Data:   "text",
		Meta:   "meta",
	}

	errSource := &model.Note{
		ID:     primitive.NewObjectID().Hex(),
		UserID: successUser.ID,
		Title:  "title",
		Data:   "text",
		Meta:   "err",
	}

	testAuth.On("GetVerifier").Return(new(jwt.Verifier))
	testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)

	app := iris.New()

	app.Post("/api/notes", testHandler.SetNote)

	app.Use(testHandler.AuthMiddleware())

	server := httptest.New(t, app, httptest.URL("http://example.com"), httptest.Debug(true), httptest.LogLevel("debug"))

	tests := []struct {
		name    string
		resCode int
		arg     *model.Note
	}{
		{
			name:    "1# success",
			resCode: http.StatusCreated,
			arg:     source,
		},
		{
			name:    "2# bad request",
			resCode: http.StatusBadRequest,
			arg:     source,
		},
		{
			name:    "3# err",
			resCode: http.StatusInternalServerError,
			arg:     errSource,
		},
		{
			name:    "4# unauthorized",
			resCode: http.StatusUnauthorized,
			arg:     source,
		},
	}

	testUsecase.On("AddNote", mock.Anything, source).Return(source, nil)
	testUsecase.On("AddNote", mock.Anything, errSource).Return(new(model.Note), errors.New("err"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case tests[0].name:
				server.POST("/api/notes").WithJSON(tt.arg).Expect().Status(tt.resCode)
			case tests[1].name:
				server.POST("/api/notes").WithJSON("foo").Expect().Status(tt.resCode)
			case tests[2].name:
				server.POST("/api/notes").WithJSON(tt.arg).Expect().Status(tt.resCode)
			case tests[3].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, errors.New("err"))
				server.POST("/api/notes").WithJSON(tt.arg).Expect().Status(tt.resCode)

			}

		})
	}
}
