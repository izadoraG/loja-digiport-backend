package main

import (
	"Lojinha-DigiPort/loja-digiport-backend/model"
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8085", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProdutos(w, r)
	}
}

func addProdutos(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	err := addProduto(produto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Erro{ErrorMessage: err.Error()})
	}

	w.WriteHeader(http.StatusCreated)
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		produtosFitradosPorNome := BuscaProdutoPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosFitradosPorNome)
	} else {
		produtos := Produtos
		json.NewEncoder(w).Encode(produtos)
	}
}
