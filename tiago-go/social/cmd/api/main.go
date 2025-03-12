package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ope/social/internal/env"
	"github.com/ope/social/internal/store"
)

func main() {
	// Get the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":9090"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
