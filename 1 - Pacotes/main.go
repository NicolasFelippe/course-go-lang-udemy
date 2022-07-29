package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

/*
Função com letra maiscula Main = publica
Função com letra minuscula main = privado

*/

func main() {
	fmt.Println("teste")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("nickasda123@gmail.com")
	fmt.Println(erro)
}
