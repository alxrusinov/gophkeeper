package middleware

import (
	"errors"
	"net/http"
	"testing"

	mockauth "github.com/alxrusinov/gophkeeper/internal/auth/mock"
	"github.com/alxrusinov/gophkeeper/internal/model"
	mockusecase "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHttpHandler_VerifyMiddleware(t *testing.T) {
	usecase := mockusecase.NewUsecaseMock()
	auth := mockauth.NewAuthMock()
	mw := NewMiddleware(usecase, auth, "user_token")

	app := iris.New()

	app.Use(mw.VerifyMiddleware)
	app.Get("/foo", func(ctx iris.Context) {
		ctx.StatusCode(http.StatusOK)
	})

	server := httptest.New(t, app, httptest.URL("http://example.com"), httptest.Debug(true), httptest.LogLevel("debug"))

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}

	auth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)

	usecase.On("CheckUser", mock.Anything, successUser.ID).Return(true, nil)

	tests := []struct {
		name    string
		userID  string
		resCode int
	}{
		{
			name:    "1# success",
			userID:  successUser.ID,
			resCode: http.StatusOK,
		},
		{
			name:    "2# unauth",
			userID:  successUser.ID,
			resCode: http.StatusUnauthorized,
		},
		{
			name:    "3# fiald check user",
			userID:  successUser.ID,
			resCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case tests[0].name:
				server.GET("/foo").Expect().Status(tt.resCode)
			case tests[1].name:
				auth.On("GetUserFromContext", mock.Anything).Unset()
				auth.On("GetUserFromContext", mock.Anything).Return(successUser, errors.New("err"))
				server.GET("/foo").Expect().Status(tt.resCode)
			case tests[2].name:
				auth.On("GetUserFromContext", mock.Anything).Unset()
				auth.On("GetUserFromContext", mock.Anything).Return(successUser, nil)

				usecase.On("CheckUser", mock.Anything, successUser.ID).Unset()
				usecase.On("CheckUser", mock.Anything, successUser.ID).Return(false, errors.New("err"))

				server.GET("/foo").Expect().Status(tt.resCode)

			}
		})
	}
}
