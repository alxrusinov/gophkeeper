package auth

import (
	"github.com/kataras/iris/v12/middleware/jwt"
)

// Auth - authenticator with jwt token
type Auth struct {
	accessSigner   *jwt.Signer
	refreshSigner  *jwt.Signer
	accessVerifier *jwt.Verifier
}

// NewAuth - new instance of Auth
func NewAuth() *Auth {
	auth := &Auth{
		accessSigner:   jwt.NewSigner(jwt.HS256, accessSecret, accessExpire).WithEncryption([]byte(accessEncriptionKey), nil),
		refreshSigner:  jwt.NewSigner(jwt.HS256, refreshSecret, refreshExpire).WithEncryption([]byte(refreshEncriptionKey), nil),
		accessVerifier: jwt.NewVerifier(jwt.HS256, accessSecret).WithDecryption([]byte(accessEncriptionKey), nil),
	}

	return auth
}
