package routes

import (
	"Lojinha-DigiPort/loja-digiport-backend/controller"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequest() {
	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.CriaProdutosHandler).Methods("POST")
	route.HandleFunc("/produtos/{id}", controller.RemoveProdutoHandler).Methods("DELETE")

	route.HandleFunc("/usuarios", controller.CriaUsuariosHandler).Methods("POST")

	http.ListenAndServe(":8080", route)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(route)

	http.ListenAndServe(":8080", handler)

}
