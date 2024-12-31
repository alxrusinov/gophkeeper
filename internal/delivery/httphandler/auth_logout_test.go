package httphandler

import (
	"net/http"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/auth"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestHttpHandler_Logout(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	testAuth := auth.NewAuth()
	testHandler := NewHttpHandler(testUsecase, testAuth)

	app := iris.New()

	app.Post("/auth/logout", testHandler.Logout)

	server := httptest.New(t, app, httptest.URL("http://example.com"))

	tests := []struct {
		name    string
		resCode int
	}{
		{
			name:    "1# success",
			resCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.POST("/auth/logout").Expect().Status(tt.resCode)
		})
	}
}
