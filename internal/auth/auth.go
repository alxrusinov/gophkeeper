package auth

import (
	"github.com/alxrusinov/gophkeeper/internal/config"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// Auth - authenticator with jwt token
type Auth struct {
	accessSigner   *jwt.Signer
	refreshSigner  *jwt.Signer
	accessVerifier *jwt.Verifier
	config         config.Config
}

// NewAuth - new instance of Auth
func NewAuth(cfg config.Config) *Auth {
	auth := &Auth{
		accessSigner:   jwt.NewSigner(jwt.HS256, cfg.Auth.AccessSecret, cfg.Auth.AccessExpire).WithEncryption([]byte(cfg.Auth.AccessEncriptionKey), nil),
		refreshSigner:  jwt.NewSigner(jwt.HS256, cfg.Auth.RefreshSecret, cfg.Auth.RefreshExpire).WithEncryption([]byte(cfg.Auth.RefreshEncriptionKey), nil),
		accessVerifier: jwt.NewVerifier(jwt.HS256, cfg.Auth.AccessSecret).WithDecryption([]byte(cfg.Auth.AccessEncriptionKey), nil),
		config:         cfg,
	}

	return auth
}
