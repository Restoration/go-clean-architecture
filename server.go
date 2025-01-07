package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"go-clean-app/config"
	"go-clean-app/infrastructure/driver"
	"go-clean-app/infrastructure/router"
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
	router.App(r, db)
	r.Run()
}
