package router

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/delivery/httphandler"
	"github.com/kataras/iris/v12"
)

// Config - interface for routers config
type Config interface {
	GetBaseURL() string
	GetFileSize() int64
}

// Router - router for server
type Router struct {
	app       *iris.Application
	handler   *httphandler.HttpHandler
	baseURL   string
	fileSizeB int64 // linit of file size in bytes
}

// Init - initialize  router
func (r *Router) init() error {

	r.app.UseGlobal(iris.Compression)
	authRouter := r.app.Party(httphandler.AuthRouteGroup)
	authRouter.Use(r.handler.Middleware.CorsMiddleware)
	authRouter.AllowMethods(iris.MethodOptions)

	authRouter.Post(httphandler.RegisterRoute, r.handler.AuthHandler.Register)
	authRouter.Post(httphandler.LoginRoute, r.handler.AuthHandler.Login)
	authRouter.Post(httphandler.LogoutRoute, r.handler.AuthHandler.Logout)

	apiRouter := r.app.Party(httphandler.ApiRouteGroup)

	apiRouter.AllowMethods(iris.MethodOptions)
	apiRouter.Use(r.handler.Middleware.CorsMiddleware)
	apiRouter.Use(r.handler.Middleware.AuthMiddleware())
	apiRouter.Use(r.handler.Middleware.VerifyMiddleware)
	apiRouter.Use(r.handler.Middleware.BodyLimitMiddleware(r.fileSizeB))

	apiRouter.Get(httphandler.NotesRoute, r.handler.NoteHandler.GetNoteList)
	apiRouter.Get(httphandler.BinariesRoute, r.handler.BinaryHandler.GetBinaryList)
	apiRouter.Get(httphandler.BankcardsRoute, r.handler.BankCardHandler.GetBankCardList)
	apiRouter.Get(httphandler.CredentialsRoute, r.handler.CredentialsHandler.GetCredentialsList)

	apiRouter.Post(httphandler.NotesRoute, r.handler.NoteHandler.SetNote)
	apiRouter.Post(httphandler.BinariesRoute, r.handler.BinaryHandler.SetBinary)
	apiRouter.Post(httphandler.BankcardsRoute, r.handler.BankCardHandler.SetBankCard)
	apiRouter.Post(httphandler.CredentialsRoute, r.handler.CredentialsHandler.SetCredentials)

	apiRouter.Get(httphandler.NoteRoute, r.handler.NoteHandler.GetNote)
	apiRouter.Get(httphandler.BinaryRoute, r.handler.BinaryHandler.GetBinary)
	apiRouter.Get(httphandler.BankcardRoute, r.handler.BankCardHandler.GetBankCard)
	apiRouter.Get(httphandler.CredentialRoute, r.handler.CredentialsHandler.GetCredentials)

	apiRouter.Delete(httphandler.NotesRoute, r.handler.NoteHandler.DeleteNote)
	apiRouter.Delete(httphandler.BinariesRoute, r.handler.BinaryHandler.DeleteBinary)
	apiRouter.Delete(httphandler.BankcardsRoute, r.handler.BankCardHandler.DeleteBankCard)
	apiRouter.Delete(httphandler.CredentialsRoute, r.handler.CredentialsHandler.DeleteCredentials)

	apiRouter.Get(httphandler.DownloadFile, r.handler.BinaryHandler.DownloadFile)

	return nil
}

// Run - run server
func (r *Router) Run(ctx context.Context) (err error) {
	err = r.init()

	if err != nil {
		return
	}

	err = r.app.Listen(r.baseURL)

	return err
}

// NewRouter - create new instance of Router
func NewRouter(cfg Config, handler *httphandler.HttpHandler) *Router {

	router := &Router{
		app:       iris.New(),
		handler:   handler,
		baseURL:   cfg.GetBaseURL(),
		fileSizeB: cfg.GetFileSize(),
	}

	return router
}
