package auth

import "github.com/kataras/iris/v12/middleware/jwt"

// GetVerifier - return custom jwt verifier
func (a *Auth) GetVerifier() *jwt.Verifier {
	return a.accessVerifier
}
