package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Create a struct App with Router and DB as fields
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Initialize() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPass, DBName)
	// Open mysql connection
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	checkError(err)

	// Initialize the router and handle routes
	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil
}

// Run the server
func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

// sendResponse function to send response
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// sendError function to send error
func sendError(w http.ResponseWriter, statusCode int, err string) {
	errorMessage := map[string]string{"error": err}
	sendResponse(w, statusCode, errorMessage)
}

// getProducts function to get products
func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}

// Handle the route
func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProduct).Methods("GET")
}
