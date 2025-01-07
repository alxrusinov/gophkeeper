package httphandler

import (
	"net/http"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) CorsMiddleware(ctx iris.Context) {

	origin := ctx.GetHeader("Origin")

	crs := cors.New(cors.Options{
		AllowedOrigins:     []string{origin},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedMethods:     []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodPut},
		MaxAge:             86400,
		Debug:              false,
		AllowedHeaders:     []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	crs(ctx)

}
