package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kmdavidds/abdimasa-backend/internal/app/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env variables: %v", err)
	}
}

func main() {
	db := config.NewDatabase()

	err := config.MigrateTables(db)
	if err != nil {
		log.Fatalf("failed to migrate tables %v", err)
	}
	
	app := config.NewFiber()

	config.StartApp(&config.AppConfig{
		DB: db,
		App: app,
	})

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}