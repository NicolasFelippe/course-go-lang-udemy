package main

import (
	"fmt"
	"os"
)

func closure() func() {
	texto := "dentro da funcao closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")

	text := "Dentro da funcao main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()
}
