package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ope/social/internal/db"
	"github.com/ope/social/internal/env"
	"github.com/ope/social/internal/store"
)

const version = "0.0.1"

func main() {
	// Get the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":9090"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:adminuser@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "dev"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	// close the db connection
	defer db.Close()
	log.Println("Database connection pool established")
	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
