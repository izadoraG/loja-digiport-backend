package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"Lojinha-DigiPort/loja-digiport-backend/model"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)
}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaProdutoPorNome(nome)
	json.NewEncoder(w).Encode(&produto)

}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	error := model.CriaProduto(produto)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

func RemoveProdutoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := model.RemoveProduto(id)
	if err != nil {
		userError := model.Erro{ErrorMessage: "ocorreu um erro ao tentar excluir o produto"}
		json.NewEncoder(w).Encode(userError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

	//model.RemoveProduto("")
}
