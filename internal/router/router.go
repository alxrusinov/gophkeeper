package router

import "github.com/kataras/iris/v12"

// Router - router for server
type Router struct {
	app *iris.Application
}

// NewRouter - create new instance of Router
func NewRouter() *Router {
	router := &Router{
		app: iris.New(),
	}

	return router
}
