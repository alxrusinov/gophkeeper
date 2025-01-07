package httphandler

import (
	"net/http"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func (h *HttpHandler) CorsMiddleware(ctx iris.Context) {
	// var (
	// 	corsOrigin             = ctx.GetHeader("Origin")
	// 	corsMethod             = ctx.GetHeader("Access-Control-Request-Method")
	// 	corsHeaders            = ctx.GetHeader("Access-Control-Request-Headers")
	// 	isCorsPreflightRequest = (corsOrigin != "") || (corsMethod != "") || (corsHeaders != "")
	// )

	// if isCorsPreflightRequest {
	// 	ctx.Header("Access-Control-Allow-Origin", corsOrigin)
	// 	ctx.Header("Access-Control-Allow-Methods", "POST, PUT, PATCH, DELETE")
	// 	ctx.Header("Access-Control-Allow-Headers", corsHeaders)
	// 	ctx.Header("Access-Control-Max-Age", "86400")
	// 	ctx.Header("Access-Control-Allow-Credentials", "true")
	// 	ctx.Header("Vary", "Access-Control-Request-Method")
	// 	ctx.Header("Access-Control-Request-Headers", "Accept,content-type,X-Requested-With,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,token")

	// 	if ctx.Request().Method == "OPTIONS" {
	// 		ctx.StatusCode(200)
	// 		return
	// 	}
	// }

	// ctx.Next()

	origin := ctx.GetHeader("Origin")

	crs := cors.New(cors.Options{
		AllowedOrigins:     []string{origin},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedMethods:     []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodPut},
		MaxAge:             86400,
		Debug:              true,
		AllowedHeaders:     []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	crs(ctx)

}
