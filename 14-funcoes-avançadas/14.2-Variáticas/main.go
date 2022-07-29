package main

import (
	"fmt"
	"os"
)

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")

	soma := soma(1, 2, 3, 4, 5, 6)
	fmt.Println("soma::::", soma)
}
