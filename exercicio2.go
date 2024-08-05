package main

import (
	"fmt"
)

func soma(arr []int) []int {
	var n []int
	for _, num := range arr {
		if num > 0 {
			n = append(n, num+10)
		} else if num == 0 {
			n = append(n, num+2)
		} else {
			n = append(n, num+23)
		}

	}
	return n
}

func mainn() {
	arr := []int{1, 2, 0, 4, -40}
	result := soma(arr)
	fmt.Println(result)
}
