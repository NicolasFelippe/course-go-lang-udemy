package main

import (
	"fmt"
	"os"
)

func recuperarExecucao() {
	fmt.Println("tentativa de recuperar a execução")
	if r := recover(); r != nil {
		fmt.Println("execução recupera com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6!")
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")
	fmt.Println(alunoEstaAprovado(6, 6))

}
