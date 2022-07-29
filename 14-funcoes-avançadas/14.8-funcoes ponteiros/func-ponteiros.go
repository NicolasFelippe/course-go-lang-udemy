package main

import (
	"fmt"
	"os"
)

func inverte(n1 int) int {
	return n1 * -1
}

func invertePOnteiro(n1 *int) {
	*n1 = *n1 * -1
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")
	numero := 20
	numeroInvertido := inverte(numero)
	fmt.Println("numeroInvertido", numeroInvertido)
	fmt.Println("numero:", numeroInvertido)

	novoNumero := 40
	fmt.Println("novoNumero:", novoNumero)
	invertePOnteiro(&novoNumero)
	fmt.Println("novoNumero:", novoNumero)
}
