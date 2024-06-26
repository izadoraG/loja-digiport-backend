package main

import (
	"fmt"
	"strings"
)

func main() {
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
