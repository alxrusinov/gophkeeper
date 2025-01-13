package middleware

import (
	"github.com/alxrusinov/gophkeeper/internal/auth"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func extractTokenFromCookie(userCookie string) jwt.TokenExtractor {
	return func(ctx iris.Context) string {
		token := ctx.GetCookie(userCookie)

		return token
	}
}

// AuthMiddleware - checks user authorization
func (m *Middleware) AuthMiddleware() context.Handler {
	verifier := m.auth.GetVerifier()

	verifier.Extractors = []jwt.TokenExtractor{extractTokenFromCookie(m.userCookie)}

	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(auth.Claims)
	})

	return verifyMiddleware
}
