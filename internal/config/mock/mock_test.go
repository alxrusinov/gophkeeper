package mock

import (
	"errors"
	"reflect"
	"testing"
)

func TestConfigMock_Run(t *testing.T) {
	cfg := NewConfigMock()
	tests := []struct {
		name    string
		cm      *ConfigMock
		wantErr bool
	}{
		{
			name:    "1# success",
			cm:      cfg,
			wantErr: false,
		},
		{
			name:    "2# error",
			cm:      cfg,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.On("Run").Unset()

			if tt.wantErr {
				cfg.On("Run").Return(errors.New("err"))
			} else {

				cfg.On("Run").Return(nil)
			}

			if err := tt.cm.Run(); (err != nil) != tt.wantErr {
				t.Errorf("ConfigMock.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigMock_GetBaseURL(t *testing.T) {
	cfg := NewConfigMock()
	result := "string"
	tests := []struct {
		name string
		cm   *ConfigMock
		want string
	}{
		{
			name: "1# success",
			cm:   cfg,
			want: result,
		},
	}
	cfg.On("GetBaseURL").Return(result)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cm.GetBaseURL(); got != tt.want {
				t.Errorf("ConfigMock.GetBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigMock_GetDbURL(t *testing.T) {
	cfg := NewConfigMock()
	result := "string"
	tests := []struct {
		name string
		cm   *ConfigMock
		want string
	}{
		{
			name: "1# success",
			cm:   cfg,
			want: result,
		},
	}

	cfg.On("GetDbURL").Return(result)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cm.GetDbURL(); got != tt.want {
				t.Errorf("ConfigMock.GetDbURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfigMock(t *testing.T) {
	result := new(ConfigMock)
	tests := []struct {
		name string
		want *ConfigMock
	}{
		{
			name: "1# success",
			want: result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigMock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigMock() = %v, want %v", got, tt.want)
			}
		})
	}
}
