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
	authRouter := r.app.Party(httphandler.AuthRouteGroup)

	authRouter.Post(httphandler.RegisterRoute, r.handler.Register)
	authRouter.Post(httphandler.LoginRoute, r.handler.Login)
	authRouter.Post(httphandler.LogoutRoute, r.handler.Logout)

	apiRouter := r.app.Party(httphandler.ApiRouteGroup)

	apiRouter.Use(r.handler.AuthMiddleware())

	apiRouter.Use(iris.Compression)

	apiRouter.Get(httphandler.NotesRoute, r.handler.GetNoteList)
	apiRouter.Get(httphandler.BinariesRoute, r.handler.GetBinaryList)
	apiRouter.Get(httphandler.BankcardsRoute, r.handler.GetBankCardList)
	apiRouter.Get(httphandler.CredentialsRoute, r.handler.GetCredentialsList)

	apiRouter.Post(httphandler.NotesRoute, r.handler.SetNote)
	apiRouter.Post(httphandler.BinariesRoute, r.handler.SetBinary)
	apiRouter.Post(httphandler.BankcardsRoute, r.handler.SetBankCard)
	apiRouter.Post(httphandler.CredentialsRoute, r.handler.SetCredentials)

	apiRouter.Get(httphandler.NoteRoute, r.handler.GetNote)
	apiRouter.Get(httphandler.BinaryRoute, r.handler.GetBinary)
	apiRouter.Get(httphandler.BankcardRoute, r.handler.GetBankCard)
	apiRouter.Get(httphandler.CredentialRoute, r.handler.GetCredentials)

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
