package auth

import (
	"reflect"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/config"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func TestAuth_GetVerifier(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Run()
	testAuth := NewAuth(*cfg)

	tests := []struct {
		name string
		a    *Auth
		want *jwt.Verifier
	}{
		{
			name: "1# success",
			a:    testAuth,
			want: testAuth.accessVerifier,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetVerifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth.GetVerifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
