package exercicios

import (
	"fmt"

	"github.com/nhanderu/brdoc"
)

func ValidatorCPF() {
	// Solicitando ao usuário que informe um CPF
	var cpf string
	fmt.Print("Informe um CPF para verificar se está na lista: ")
	fmt.Scanln(&cpf)

	// Validando o CPF
	if brdoc.IsCPF(cpf) {
		fmt.Println("Este é um CPF válido:", cpf)
	} else {
		fmt.Println("Este não é um CPF válido:", cpf)
	}
}
