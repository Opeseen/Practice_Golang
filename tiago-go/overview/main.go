package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":9090"}

	// Initialize the ServeMux
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
