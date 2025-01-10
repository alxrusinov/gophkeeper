package auth

import (
	"strings"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/config"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAuth_GetTokenPair(t *testing.T) {
	cfg := config.NewConfig()
	cfg.RunMock()
	auth := NewAuth(*cfg)
	user := &model.User{
		Username: "foo",
		ID:       "111",
	}

	access, _ := auth.GetAccessToken(user)
	refresh, _ := auth.GetRefreshToken(user)

	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		a       *Auth
		args    args
		want    *model.TokenPair
		wantErr bool
	}{
		{
			name: "1# success",
			a:    auth,
			args: args{
				user: user,
			},
			want: &model.TokenPair{
				AccessToken:  access,
				RefreshToken: refresh,
				Exp:          int64(cfg.Auth.AccessExpire.Seconds()),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetTokenPair(tt.args.user)

			assert.Equal(t, strings.Split(tt.want.AccessToken, ".")[0], strings.Split(got.AccessToken, ".")[0])
			assert.Equal(t, strings.Split(tt.want.RefreshToken, ".")[0], strings.Split(got.RefreshToken, ".")[0])

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}
