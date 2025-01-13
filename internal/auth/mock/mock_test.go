package mock

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAuthMock_GetAccessTokenExp(t *testing.T) {
	auth := NewAuthMock()
	result := time.Hour * 24

	tests := []struct {
		name string
		am   *AuthMock
		want time.Duration
	}{
		{
			name: "1# success",
			am:   auth,
			want: result,
		},
	}

	auth.On("GetAccessTokenExp").Return(result)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.am.GetAccessTokenExp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthMock.GetAccessTokenExp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthMock_GetAccessToken(t *testing.T) {
	auth := NewAuthMock()
	result := "token"

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}

	errUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}

	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		am      *AuthMock
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1# success",
			am:   auth,
			args: args{
				user: successUser,
			},
			want:    result,
			wantErr: false,
		},
		{
			name: "2# error",
			am:   auth,
			args: args{
				user: errUser,
			},
			want:    "",
			wantErr: true,
		},
	}

	auth.On("GetAccessToken", successUser).Return(result, nil)
	auth.On("GetAccessToken", errUser).Return("", errors.New("err"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.am.GetAccessToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthMock.GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthMock.GetAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthMock_GetRefreshToken(t *testing.T) {
	auth := NewAuthMock()
	result := "token"

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}

	errUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		am      *AuthMock
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1# success",
			am:   auth,
			args: args{
				user: successUser,
			},
			want:    result,
			wantErr: false,
		},
		{
			name: "2# error",
			am:   auth,
			args: args{
				user: errUser,
			},
			want:    "",
			wantErr: true,
		},
	}
	auth.On("GetRefreshToken", successUser).Return(result, nil)
	auth.On("GetRefreshToken", errUser).Return("", errors.New("err"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.am.GetRefreshToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthMock.GetRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthMock.GetRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthMock_GetTokenPair(t *testing.T) {
	auth := NewAuthMock()
	result := &model.TokenPair{
		AccessToken:  "access",
		RefreshToken: "refresh",
		Exp:          123,
	}

	successUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}

	errUser := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}

	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		am      *AuthMock
		args    args
		want    *model.TokenPair
		wantErr bool
	}{
		{
			name: "1# success",
			am:   auth,
			args: args{
				user: successUser,
			},
			want:    result,
			wantErr: false,
		},
		{
			name: "2# error",
			am:   auth,
			args: args{
				user: errUser,
			},
			want:    result,
			wantErr: true,
		},
	}
	auth.On("GetTokenPair", successUser).Return(result, nil)
	auth.On("GetTokenPair", errUser).Return(result, errors.New("err"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.am.GetTokenPair(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthMock.GetTokenPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthMock.GetTokenPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthMock_GetVerifier(t *testing.T) {
	auth := NewAuthMock()
	result := new(jwt.Verifier)

	tests := []struct {
		name string
		am   *AuthMock
		want *jwt.Verifier
	}{
		{
			name: "1# success",
			am:   auth,
			want: result,
		},
	}
	auth.On("GetVerifier").Return(result)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.am.GetVerifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthMock.GetVerifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthMock_RefreshUserTokens(t *testing.T) {
	auth := NewAuthMock()
	result := &model.TokenPair{
		AccessToken:  "access",
		RefreshToken: "refresh",
		Exp:          123,
	}

	var successCtx iris.Context
	var errCtx iris.Context

	type args struct {
		ctx iris.Context
	}
	tests := []struct {
		name    string
		am      *AuthMock
		args    args
		want    *model.TokenPair
		wantErr bool
	}{
		{
			name: "1# success",
			am:   auth,
			args: args{
				ctx: successCtx,
			},
			want:    result,
			wantErr: false,
		},
		{
			name: "2# error",
			am:   auth,
			args: args{
				ctx: errCtx,
			},
			want:    result,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth.On("RefreshUserTokens", successCtx).Unset()

			if tt.wantErr {
				auth.On("RefreshUserTokens", errCtx).Return(result, errors.New("err"))
			} else {
				auth.On("RefreshUserTokens", successCtx).Return(result, nil)
			}

			got, err := tt.am.RefreshUserTokens(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthMock.RefreshUserTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthMock.RefreshUserTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthMock_GetUserFromContext(t *testing.T) {
	auth := NewAuthMock()
	result := &model.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: "user",
	}

	var successCtx iris.Context
	var errCtx iris.Context

	type args struct {
		ctx iris.Context
	}
	tests := []struct {
		name    string
		am      *AuthMock
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "1# success",
			am:   auth,
			args: args{
				ctx: successCtx,
			},
			want:    result,
			wantErr: false,
		},
		{
			name: "2# error",
			am:   auth,
			args: args{
				ctx: errCtx,
			},
			want:    result,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth.On("GetUserFromContext", successCtx).Unset()

			if tt.wantErr {
				auth.On("GetUserFromContext", errCtx).Return(result, errors.New("err"))
			} else {
				auth.On("GetUserFromContext", successCtx).Return(result, nil)
			}
			got, err := tt.am.GetUserFromContext(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthMock.GetUserFromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthMock.GetUserFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
