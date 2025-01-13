package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/alxrusinov/gophkeeper/internal/app"
)

func main() {

	application := app.NewApp()

	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	defer stop()

	if err := application.Run(signalCtx); err != nil {
		log.Fatal("application was crashed", err)
	}
}
