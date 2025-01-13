package customerrors

import (
	"errors"
	"testing"
)

func TestNotFound_Unwrap(t *testing.T) {
	result := errors.New("error")

	tests := []struct {
		name    string
		err     *NotFound
		wantErr bool
	}{
		{
			name: "1# success",
			err: &NotFound{
				Err: result,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.err.Unwrap(); (err != nil) != tt.wantErr {
				t.Errorf("NotFound.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotFound_Error(t *testing.T) {
	result := errors.New("error")

	tests := []struct {
		name string
		err  *NotFound
		want string
	}{
		{
			name: "1# success",
			err: &NotFound{
				Err: result,
			},
			want: result.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("NotFound.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
