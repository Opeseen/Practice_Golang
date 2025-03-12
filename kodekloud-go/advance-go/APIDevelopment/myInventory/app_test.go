package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialize(DBUser, DBPass, DBName)
	if err != nil {
		log.Fatal("Error occurred while initializing the database")
	}
	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS test_product(
		id int NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		quantity int,
		price float(10, 2),
		PRIMARY KEY (id)
	);`

	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM test_product")
	a.DB.Exec("ALTER TABLE test_product AUTO_INCREMENT=1")
	log.Println("clear table")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT INTO test_product(name, quantity, price) VALUES('%v', '%v', '%v')", name, quantity, price)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.Println(err)
	}

}

func sentRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)

	return recorder
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v Received %v", expectedStatusCode, actualStatusCode)
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 100, 50)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sentRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)
}

func TestCreateProduct(t *testing.T) {
	clearTable()
	var product = []byte(`{"name":"chair", "quantity": 1, "price": 100}`)
	request, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	request.Header.Set("Content-Type", "application/json")

	response := sentRequest(request)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]any
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "chair" {
		t.Errorf("Expected name: %v, Got: %v", "chair", m["name"])
	}

	if m["quantity"] != 1.0 {
		t.Errorf("Expected quantity: %v, Got: %v", 1.0, m["quantity"])
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProduct("connector", 10, 10)

	request, _ := http.NewRequest("GET", "/product/6", nil)
	response := sentRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("DELETE", "/product/6", nil)
	response = sentRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("GET", "/product/6", nil)
	response = sentRequest(request)
	checkStatusCode(t, http.StatusNotFound, response.Code)
}
