package bankcardhandler

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	authMock "github.com/alxrusinov/gophkeeper/internal/auth/mock"
	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler/middleware"
	"github.com/alxrusinov/gophkeeper/internal/model"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHttpHandler_GetBankCard(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	testAuth := authMock.NewAuthMock()
	testHandler := NewBankCardHandler(testUsecase, testAuth)
	mw := middleware.NewMiddleware(testUsecase, testAuth, "user_token")

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

	app.Get("/api/bankcard/{id}", testHandler.GetBankCard)

	app.Use(mw.AuthMiddleware())

	server := httptest.New(t, app, httptest.URL("http://example.com"), httptest.Debug(true), httptest.LogLevel("debug"))

	errAuth := authMock.NewAuthMock()

	errHandler := NewBankCardHandler(testUsecase, errAuth)

	errApp := iris.New()

	errApp.Get("/api/bankcard/{id}", errHandler.GetBankCard)

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

	testUsecase.On("GetBankCard", mock.Anything, mock.Anything, successNoteID).Return(&model.BankCard{
		ID:     successNoteID,
		UserID: successUser.ID,
		Title:  "Title",
		Data:   1233,
		Meta:   "Meta info",
	}, nil)
	testUsecase.On("GetBankCard", mock.Anything, mock.Anything, errID).Return(new(model.BankCard), errors.New("err"))
	testUsecase.On("GetBankCard", mock.Anything, mock.Anything, noNote).Return(new(model.BankCard), &customerrors.NotFound{Err: errors.New("err")})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case tests[0].name:
				server.GET(fmt.Sprintf("/api/bankcard/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[1].name:
				server.GET(fmt.Sprintf("/api/bankcard/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[2].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(new(model.User), errors.New("error of context"))
				errServer.GET(fmt.Sprintf("/api/bankcard/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[3].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)
				server.GET(fmt.Sprintf("/api/bankcard/%s", tt.arg)).Expect().Status(tt.resCode)
			case tests[4].name:
				testAuth.On("GetUserFromContext", mock.Anything).Unset()
				testAuth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)
				server.GET(fmt.Sprintf("/api/bankcard/%s", tt.arg)).Expect().Status(tt.resCode)

			}

		})
	}
}
