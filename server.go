package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	"go-clean-app/config"
	"go-clean-app/infrastructure/driver"
	"go-clean-app/infrastructure/router"
	"go-clean-app/infrastructure/tracer"
)

func main() {
	r := gin.Default()
	err := config.Load("")
	if err != nil {
		log.Fatalf("Failed to load env: %v\n", err)
	}
	db := driver.Initialize()
	db.CloseConnections()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	otelShutdown, err := tracer.SetupOTelSDK(ctx)
	if err != nil {
		log.Fatalf("error in open telemetry: %v\n", err)
	}
	defer func() {
		if err := errors.Join(err, otelShutdown(context.Background())); err != nil {
			log.Fatalf("open telemetry shutdown error: %v\n", err)
		}
	}()

	router.App(r, db)
	r.Run()
}
