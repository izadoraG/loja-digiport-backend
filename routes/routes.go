package routes

import (
	"Lojinha-DigiPort/loja-digiport-backend/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.CriaProdutosHandler).Methods("POST")
	http.ListenAndServe(":8080", route)
}
