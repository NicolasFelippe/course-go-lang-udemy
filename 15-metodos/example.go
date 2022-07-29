package main

import (
	"fmt"
	"os"
)

type usuario struct {
	name  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("salvando usuario")
}

func (u *usuario) alterarCampo() {
	u.idade++
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")
	user1 := usuario{"teste", 20}
	fmt.Println(user1)
	user1.alterarCampo()
}
