package auth

import (
	"reflect"
	"testing"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/config"
)

func TestAuth_GetAccessTokenExp(t *testing.T) {
	cfg := config.NewConfig()
	cfg.RunMock()
	testAuth := NewAuth(*cfg)

	tests := []struct {
		name string
		a    *Auth
		want time.Duration
	}{
		{
			name: "1# success",
			a:    testAuth,
			want: testAuth.config.Auth.AccessExpire,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetAccessTokenExp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth.GetAccessTokenExp() = %v, want %v", got, tt.want)
			}
		})
	}
}
