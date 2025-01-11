package router

import (
	"context"
	"testing"

	authmock "github.com/alxrusinov/gophkeeper/internal/auth/mock"
	"github.com/alxrusinov/gophkeeper/internal/config"

	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler"
	usecasemock "github.com/alxrusinov/gophkeeper/internal/usecase/mock"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func TestRouter_init(t *testing.T) {
	cfg := config.NewConfig()
	cfg.RunMock()

	uc := usecasemock.NewUsecaseMock()
	am := authmock.NewAuthMock()
	am.On("GetVerifier").Return(new(jwt.Verifier))
	handler := httphandler.NewHttpHandler(uc, am)
	router := NewRouter(cfg, handler)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		r       *Router
		args    args
		wantErr bool
	}{
		{
			name: "1# success",
			r:    router,
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.init(); (err != nil) != tt.wantErr {
				t.Errorf("Router.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
