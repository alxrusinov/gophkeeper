package auth

import (
	"reflect"
	"testing"
	"time"
)

func TestAuth_GetAccessTokenExp(t *testing.T) {
	tests := []struct {
		name string
		a    *Auth
		want time.Duration
	}{
		{
			name: "1# success",
			a:    NewAuth(),
			want: accessExpire,
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
