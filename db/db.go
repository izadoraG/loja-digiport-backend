package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	dbPass := os.Getenv("DB_PASS")
	connStr := fmt.Sprint("user=postgres dbname=digiport_loja password=", dbPass, " host=localhost sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
