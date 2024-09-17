package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Lojinha-DigiPort/loja-digiport-backend/model"
)

func CriaUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	fmt.Println(usuario.Email)

	err := model.CriaUsuario(usuario)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
