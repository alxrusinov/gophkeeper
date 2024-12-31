package customerrors

import (
	"errors"
	"testing"
)

func TestUnverifiedUser_Unwrap(t *testing.T) {
	result := errors.New("error")

	tests := []struct {
		name    string
		err     *UnverifiedUser
		wantErr bool
	}{
		{
			name: "1# success",
			err: &UnverifiedUser{
				Err: result,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.err.Unwrap(); (err != nil) != tt.wantErr {
				t.Errorf("UnverifiedUser.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnverifiedUser_Error(t *testing.T) {
	result := errors.New("error")

	tests := []struct {
		name string
		err  *UnverifiedUser
		want string
	}{
		{
			name: "1# success",
			err: &UnverifiedUser{
				Err: result,
			},
			want: result.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("UnverifiedUser.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
