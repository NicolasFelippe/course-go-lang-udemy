package main

import (
	"fmt"
	"os"
)

//Retorno nomeado
func soma(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return soma, subtracao
}
func soma2(n1, n2 int) (soma int, subtracao int) {
	return n1 + n2, n1 - n2
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")

	soma, subtrcao := soma2(1, 1)
	fmt.Println("soma2::::", soma, subtrcao)
}
