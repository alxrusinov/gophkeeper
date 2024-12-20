package router

import (
	"context"

	"github.com/kataras/iris/v12"
)

// Config - interface for routers config
type Config interface {
	GetBaseURL() string
}

// Router - router for server
type Router struct {
	app     *iris.Application
	baseURL string
}

// Run - runner for router
func (r *Router) Run(ctx context.Context) (err error) {
	authRouter := r.app.Party("/auth")

	authRouter.Post("/register", r.Register)
	authRouter.Post("/login", r.Login)

	apiRouter := r.app.Party("/api")

	apiRouter.Use(iris.Compression)

	apiRouter.Get("/notes", r.GetNotes)

	err = r.app.Listen(r.baseURL)

	return err
}

// NewRouter - create new instance of Router
func NewRouter(cfg Config) *Router {
	router := &Router{
		app:     iris.New(),
		baseURL: cfg.GetBaseURL(),
	}

	return router
}
