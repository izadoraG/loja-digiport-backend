package exercicios

import (
	"fmt"
)

func exercicio2() {
	// Definindo variáveis para os gastos fixos
	var aluguel float64 = 1000.0
	var internet float64 = 100.0
	var transporte float64 = 300.0

	// Calculando o total de gastos
	totalGastos := aluguel + internet + transporte

	// Imprimindo o total de gastos
	fmt.Printf("Total de gastos mensais: R$ %.2f\n", totalGastos)

	// Definindo o orçamento máximo permitido
	orcamentoMaximo := 2000.0

	// Verificando se o total de gastos excede o orçamento máximo
	if totalGastos > orcamentoMaximo {
		fmt.Println("Os gastos mensais excederam o orçamento máximo.")
	} else {
		fmt.Println("Os gastos mensais estão dentro do orçamento.")
	}
}
