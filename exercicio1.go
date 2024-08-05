package main

import (
	"fmt"
)

func mainn() {
	var num int
	fmt.Printf("Digite um numero:")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("greater than zero")
	} else if num == 0 {
		fmt.Println("zero")
	} else {
		fmt.Println("less than zero")
	}

}
