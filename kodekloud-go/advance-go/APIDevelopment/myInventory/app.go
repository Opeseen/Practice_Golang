package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProds(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}

// getProduct function to get product
func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get the id from the request
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	p := Product{ID: key}
	err = p.getProd(app.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Product not found")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w, http.StatusOK, p)
}

// createProduct function to create product
func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = p.createProd(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusCreated, p)
}

// updateProduct function to update product
func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get the id from the request
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p Product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.ID = key

	err = p.updateProd(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, p)
}

// deleteProduct function to delete product
func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get the id from the request
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	p := Product{ID: key}
	err = p.deleteProd(app.DB)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": "successful deletion"})

}

// Handle the route
func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")
}
