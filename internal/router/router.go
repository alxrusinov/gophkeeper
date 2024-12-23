package router

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler"
	"github.com/kataras/iris/v12"
)

// Config - interface for routers config
type Config interface {
	GetBaseURL() string
}

// Router - router for server
type Router struct {
	app     *iris.Application
	handler *httphandler.HttpHandler
	baseURL string
}

// Run - runner for router
func (r *Router) Run(ctx context.Context) (err error) {
	r.app.Use(iris.Compression)
	authRouter := r.app.Party(httphandler.ApiRouteGroup)

	authRouter.Post(httphandler.RegisterRoute, r.handler.Register)
	authRouter.Post(httphandler.LoginRoute, r.handler.Login)
	authRouter.Post(httphandler.LogoutRoute, r.handler.Logout)

	apiRouter := r.app.Party(httphandler.ApiRouteGroup)

	apiRouter.Use(r.handler.AuthMiddleware())

	apiRouter.Use(iris.Compression)

	apiRouter.Get(httphandler.NotesRoute, r.handler.GetNotes)
	apiRouter.Get(httphandler.BinariesRoute, r.handler.GetBinaries)
	apiRouter.Get(httphandler.BankcardsRoute, r.handler.GetBankCards)
	apiRouter.Get(httphandler.CredentialsRoute, r.handler.GetCredentials)

	apiRouter.Post(httphandler.NoteRoute, r.handler.GetNote)
	apiRouter.Post(httphandler.BinaryRoute, r.handler.GetBinary)
	apiRouter.Post(httphandler.BankcardRoute, r.handler.GetBankCard)
	apiRouter.Post(httphandler.CredentialRoute, r.handler.GetCredential)

	apiRouter.Delete(httphandler.NotesRoute, r.handler.DeleteNote)
	apiRouter.Delete(httphandler.BinariesRoute, r.handler.DeleteBinary)
	apiRouter.Delete(httphandler.BankcardsRoute, r.handler.DeleteBankCard)
	apiRouter.Delete(httphandler.CredentialsRoute, r.handler.DeleteCredentials)

	err = r.app.Listen(r.baseURL)

	return err
}

// NewRouter - create new instance of Router
func NewRouter(cfg Config, handler *httphandler.HttpHandler) *Router {
	router := &Router{
		app:     iris.New(),
		handler: handler,
		baseURL: cfg.GetBaseURL(),
	}

	return router
}
