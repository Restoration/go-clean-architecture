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
	defer func() {
		conn, err := db.DB()
		if err != nil {
			log.Fatalf("connection failed: %v\n", err)
		}
		if err := conn.Close(); err != nil {
			log.Fatalf("error in close database: %v\n", err)
		}
	}()

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
	tracer.RegisterGORMCallbacks(db)

	router.App(r, db)
	r.Run()
}
