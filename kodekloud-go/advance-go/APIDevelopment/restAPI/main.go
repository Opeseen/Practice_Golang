package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: homepage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Vars is: ", vars)
	key := vars["id"]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	myRouter.HandleFunc("/", homepage)
	http.ListenAndServe(":10000", myRouter)
}

func main() {
	Products = []Product{
		{Id: "1", Name: "Laptop", Quantity: 100, Price: 50000},
		{Id: "2", Name: "Mobile", Quantity: 200, Price: 20000},
		{Id: "3", Name: "Tablet", Quantity: 300, Price: 10000},
	}
	handleRequests()
}
