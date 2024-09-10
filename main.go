package main

import "Lojinha-DigiPort/loja-digiport-backend/db"

func main() {
	db.InitDB()
	StartServer()
}

/*import "fmt"

func main() {
	nome := "Blusa Floralll" // substituir por um nome de produto

	produtosFiltrados := BuscaProdutoPorNome(nome)

	fmt.Println(produtosFiltrados)
}*/
