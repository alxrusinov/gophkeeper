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
	authRouter := r.app.Party(authRouteGroup)

	authRouter.Post(registerRoute, r.Register)
	authRouter.Post(loginRoute, r.Login)

	apiRouter := r.app.Party(apiRouteGroup)

	apiRouter.Use(iris.Compression)

	apiRouter.Get(notesRoute, r.GetNotes)
	apiRouter.Get(binariesRoute, r.GetBinaries)
	apiRouter.Get(bankcardsRoute, r.GetBankCards)
	apiRouter.Get(credentialsRoute, r.GetCredentials)

	apiRouter.Post(noteRoute, r.GetNote)
	apiRouter.Post(binaryRoute, r.GetBinary)
	apiRouter.Post(bankcardRoute, r.GetBankCard)
	apiRouter.Post(credentialRoute, r.GetCredential)

	apiRouter.Delete(notesRoute, r.DeleteNote)
	apiRouter.Delete(binariesRoute, r.DeleteBinary)
	apiRouter.Delete(bankcardsRoute, r.DeleteBankCard)
	apiRouter.Delete(credentialsRoute, r.DeleteCredentials)

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
