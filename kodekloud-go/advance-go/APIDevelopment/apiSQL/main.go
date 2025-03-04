package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Data struct to store sql response
type Data struct {
	id   int
	name string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPass, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	result, err := db.Exec("INSERT into data values(4,'xyz')")
	checkError(err)
	lastInsertedId, err := result.LastInsertId()
	fmt.Println("Last inserted id is: ", lastInsertedId)
	checkError(err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("Number of rows affected: ", rowsAffected)
	checkError(err)
	rows, err := db.Query("SELECT * FROM data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}

}
