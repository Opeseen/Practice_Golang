package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Get all products from the database
func getProds(db *sql.DB) ([]Product, error) {
	query := "SELECT id, name, quantity, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	product := []Product{}
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}
		product = append(product, p)
	}
	return product, nil
}

func (p *Product) getProd(db *sql.DB) error {
	query := fmt.Sprintf("SELECT name, quantity, price FROM products WHERE id=%v", p.ID)
	row := db.QueryRow(query)
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) createProd(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO products(name, quantity, price) VALUES('%v', %v, %v)", p.Name, p.Quantity, p.Price)
	result, err := db.Exec(query)
	checkError(err)
	id, err := result.LastInsertId()
	checkError(err)
	p.ID = int(id)
	return nil
}

func (p *Product) updateProd(db *sql.DB) error {
	query := fmt.Sprintf("UPDATE products SET name='%v', quantity='%v', price='%v' WHERE id=%v ", p.Name, p.Quantity, p.Price, p.ID)
	result, err := db.Exec(query)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no changes affected")
	}
	return err

}

func (p *Product) deleteProd(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM products WHERE id=%v ", p.ID)
	result, err := db.Exec(query)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no product found to delete")
	}
	return err

}
