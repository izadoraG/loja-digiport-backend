package controller

import (
	"encoding/json"
	"net/http"

	"Lojinha-DigiPort/loja-digiport-backend/model"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)
}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaProdutoPorNome(nome)
	json.NewEncoder(w).Encode(produto)

}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {

}
