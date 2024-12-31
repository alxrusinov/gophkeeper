package logger

import "testing"

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "1# success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitLogger(); (err != nil) != tt.wantErr {
				t.Errorf("InitLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
