package main

import (
	"log"

	"github.com/ope/social/internal/db"
	"github.com/ope/social/internal/env"
	"github.com/ope/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:adminuser@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(&store, conn)
}
