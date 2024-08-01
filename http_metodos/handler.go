package main

import (
	"Lojinha-DigiPort/loja-digiport-backend/pessoa"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Rotas() *mux.Router {
	rotas := mux.NewRouter()
	rotas.HandleFunc("/Lista/pessoas", HandlePessoa).Methods("GET")
	return rotas

}

func HandlePessoa(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoa.ListaPessoa())
}
