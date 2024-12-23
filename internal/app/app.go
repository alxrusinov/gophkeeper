package app

import (
	"context"

	"github.com/alxrusinov/gophkeeper/internal/logger"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/alxrusinov/gophkeeper/internal/repository/mongo"
	"github.com/alxrusinov/gophkeeper/internal/router"
	"github.com/alxrusinov/gophkeeper/internal/usecase"
)

// App - aplication
type App struct {
	config Config
}

// Config - interface for config
type Config interface {
	model.Runner
	GetBaseURL() string
}

// Run - method of running application
func (app *App) Run(ctx context.Context) (err error) {
	err = app.config.Run()

	if err != nil {
		return err
	}

	err = logger.InitLogger()

	if err != nil {
		return err
	}

	repo := mongo.NewMongo()

	currentUsecase := usecase.NewUsecase(repo)

	server := router.NewRouter(app.config, currentUsecase)

	err = server.Run(ctx)

	return err
}

// NewApp - create new application
func NewApp(config Config) *App {
	application := &App{config: config}

	return application
}
