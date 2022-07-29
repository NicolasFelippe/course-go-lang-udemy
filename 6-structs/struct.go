package main

import "fmt"

type usuario struct {
	name     string
	age      uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("structs")

	var u usuario
	fmt.Println(u) // { 0}

	u.name = "Nicolas"
	u.age = 25
	fmt.Println(u)

	enderecou2 := endereco{"Rua x", 234}
	u2 := usuario{"Carol", 24, enderecou2}
	fmt.Println(u2)

	u3 := usuario{name: "Davi"}
	fmt.Println(u3)

	u4 := usuario{age: 22}
	fmt.Println(u4)
}
