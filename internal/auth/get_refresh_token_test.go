package auth

import (
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/config"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAuth_GetRefreshToken(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Run()
	user := &model.User{
		Username: "foo",
		ID:       "123",
	}

	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		a       *Auth
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1# success",
			a:    NewAuth(*cfg),
			args: args{
				user: user,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetRefreshToken(tt.args.user)

			if tt.name == tests[0].name {
				assert.Nil(t, err)
				assert.NotEmpty(t, got)
			}

		})
	}
}
