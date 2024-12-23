package router

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/auth"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/kataras/iris/v12"
)

// Config - interface for routers config
type Config interface {
	GetBaseURL() string
}

// Usecase - interface of Usecase
type Usecase interface {
	// VerifyUser - return information about user, if user exists
	// fact of user existing and error
	VerifyUser(lg *model.Login) (*model.User, error)
}

// Router - router for server
type Router struct {
	app     *iris.Application
	auth    *auth.Auth
	usecase Usecase
	baseURL string
}

// Run - runner for router
func (r *Router) Run(ctx context.Context) (err error) {
	r.app.Use(iris.Compression)
	authRouter := r.app.Party(authRouteGroup)

	authRouter.Post(registerRoute, r.Register)
	authRouter.Post(loginRoute, r.Login)
	authRouter.Post(logoutRoute, r.Logout)

	apiRouter := r.app.Party(apiRouteGroup)

	apiRouter.Use(r.AuthMiddleware())

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
func NewRouter(cfg Config, usecase Usecase) *Router {
	router := &Router{
		app:     iris.New(),
		auth:    auth.NewAuth(),
		usecase: usecase,
		baseURL: cfg.GetBaseURL(),
	}

	return router
}
