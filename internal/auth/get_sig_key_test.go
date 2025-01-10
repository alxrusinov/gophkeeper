package auth

import (
	"reflect"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/config"
)

func TestAuth_GetSigKey(t *testing.T) {
	cfg := config.NewConfig()
	cfg.RunMock()
	tests := []struct {
		name string
		a    *Auth
		want []byte
	}{
		{
			name: "1# success",
			a:    NewAuth(*cfg),
			want: []byte(cfg.Auth.SigKey),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetSigKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth.GetSigKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
