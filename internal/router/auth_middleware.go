package router

import (
	"github.com/alxrusinov/gophkeeper/internal/auth"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func extractTokenFromCookie(ctx iris.Context) string {
	token := ctx.GetCookie(userCookie)

	return token
}

// AuthMiddleware - checks user authorization
func (r *Router) AuthMiddleware() context.Handler {
	verifier := r.auth.GetVerifier()

	verifier.Extractors = []jwt.TokenExtractor{extractTokenFromCookie}

	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(auth.Claims)
	})

	return verifyMiddleware
}
