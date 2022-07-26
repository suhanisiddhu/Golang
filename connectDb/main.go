package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Suhani@123@tcp(127.0.0.1:3306)/suhani")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	dat, err := db.Query("SELECT * from Persons;")
	defer dat.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dat.Columns())
}
