package main

import (
	"fmt"
	"os"
)

func funcao1() {
	fmt.Println("Escrevendo funcao1")
}

func funcao2() {
	fmt.Println("Escrevendo funcao2")
}

func funcao3(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Entrnado na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")
	// defer funcao1()
	// funcao2()
	fmt.Println(funcao3(7, 8))
}
