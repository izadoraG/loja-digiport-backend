package exercicios

import (
	"fmt"
)

// Definindo a struct Contato
type Contato struct {
	Telefone string
	Email    string
}

func Exercicio3() {
	// Criando um map onde as chaves são strings (nomes) e os valores são structs de Contato
	contatos := map[string]Contato{
		"Alice": {"1234-5678", "alice@email.com"},
		"Bob":   {"9876-5432", "bob@email.com"},
		"Carol": {"1111-2222", "carol@email.com"},
	}

	// Solicitando ao usuário que informe um nome de contato para verificar
	var nome string
	fmt.Print("Informe um contato para verificar se está na lista: ")
	fmt.Scanln(&nome)

	// Verificando se o contato existe no map
	contato, encontrado := contatos[nome]
	if encontrado {
		// Imprimindo os detalhes do contato encontrado
		fmt.Println("Detalhes do Contato:")
		fmt.Println("Nome:", nome)
		fmt.Println("Telefone:", contato.Telefone)
		fmt.Println("Email:", contato.Email)
	} else {
		// Caso o contato não seja encontrado
		fmt.Printf("O contato com o nome '%s' não foi encontrado na agenda.\n", nome)
	}
}
