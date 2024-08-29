package main

import "fmt"

func main() {

}

func soma(i, j int) int {
	return i + j
}

func validaNome(nome string) error {
	if nome == "" {
		return fmt.Errorf("nome não preenchido")
	}
	return nil
}
