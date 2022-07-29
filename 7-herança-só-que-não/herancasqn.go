package main

import "fmt"

type pessoa struct {
	name string
	age  uint8
}

type estudante struct {
	pessoa    // isso é "herança", go copia os atributos da estrutura.
	curso     string
	faculdade string
}

type estudanteSemHeranca struct {
	pessoa    pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("heranca so que nao")

	p1 := pessoa{"Nicolas", 25}
	fmt.Println(p1)

	e1 := estudante{p1, "curso1", "faculdade1"}
	fmt.Println(e1)
	fmt.Println(e1.name)
	e2 := estudanteSemHeranca{p1, "curso1", "faculdade1"}
	fmt.Println(e2)
	fmt.Println(e2.pessoa.name)
}
