package main

import "fmt"

func main() {
	fmt.Println("--------Estrutura de controle-----")
	num := 10
	if num > 4 {
		fmt.Println("maior que 4")
	} else {
		fmt.Println("menor que 4")
	}

	if outroNumero := num; outroNumero > 0 {
		fmt.Println("maior que 0")
	}
}
