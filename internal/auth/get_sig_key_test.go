package auth

import (
	"reflect"
	"testing"
)

func TestAuth_GetSigKey(t *testing.T) {
	tests := []struct {
		name string
		a    *Auth
		want []byte
	}{
		{
			name: "1# success",
			a:    NewAuth(),
			want: []byte(sigKey),
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
