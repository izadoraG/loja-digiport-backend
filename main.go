package main

import (
	"fmt"
)

/*func main() {
	despesas := []string{"mercado", "aluguel", "eletricidade"}

	fmt.Println("Despesas a serem pagas:")
	for _, despesa := range despesas {
		fmt.Println("- ", despesa)
	}

	var item string
	fmt.Print("\nDigite um item para verificar se está na lista: ")
	fmt.Scanln(&item)

	encontrado := false
	for _, d := range despesas {
		if strings.ToLower(d) == strings.ToLower(item) {
			encontrado = true
			break
		}
	}


	if encontrado {
		fmt.Printf("%s está na lista de despesas.\n", item)
	} else {
		fmt.Printf("%s não está na lista de despesas.\n", item)
	}

	total := len(despesas)
	fmt.Printf("\nTotal de itens na lista de despesas: %d\n", total)
}
*/


func main() {
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
