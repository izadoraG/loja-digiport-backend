package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8085", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {

	nome := "Blusa Floral"
	produto := BuscaProdutoPorNome(nome)

	json.NewEncoder(w).Encode(produto)
}
