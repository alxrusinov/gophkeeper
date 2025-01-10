package httphandler

import (
	"net/http"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/auth"
	"github.com/alxrusinov/gophkeeper/internal/config"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestHttpHandler_Logout(t *testing.T) {
	testUsecase := usecasemock.NewUsecaseMock()
	cfg := config.NewConfig()
	cfg.RunMock()
	testAuth := auth.NewAuth(*cfg)
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
