package customerrors

import (
	"errors"
	"testing"
)

func TestNonexistentUser_Unwrap(t *testing.T) {
	result := errors.New("error")

	tests := []struct {
		name    string
		err     *NonexistentUser
		wantErr bool
	}{
		{
			name: "1# success",
			err: &NonexistentUser{
				Err: result,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.err.Unwrap(); (err != nil) != tt.wantErr {
				t.Errorf("NonexistentUser.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNonexistentUser_Error(t *testing.T) {
	result := errors.New("error")

	tests := []struct {
		name string
		err  *NonexistentUser
		want string
	}{
		{
			name: "1# success",
			err: &NonexistentUser{
				Err: result,
			},
			want: result.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("NonexistentUser.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
