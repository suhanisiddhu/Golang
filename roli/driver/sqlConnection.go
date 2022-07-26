package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connection() *sql.DB {
	Db, err := sql.Open("mysql", "root:Suhani@123@tcp(127.0.0.1:3306)/AuthorBook")
	if err != nil {
		log.Fatal("failed to connect with database:\n", err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal("failed to ping", pingErr)
	}

	return Db
}
