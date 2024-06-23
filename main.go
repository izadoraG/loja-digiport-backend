package main
import (
	"database/sql"
	"fmt"
    "log"

	_"github.com/lib/pq"
)

func main () {
	db := ConectaBancoDados()
	fmt.Println ("Bem vindo(a) a loja Digiport!")

	defer db.Close()

}

func ConectaBancoDados() *sql.DB {
	connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}