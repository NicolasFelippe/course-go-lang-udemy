package main

import (
	"fmt"
)

func main() {
	fmt.Println("LOOPS -------")
	i := 0
	for i < 10 {
		i++
		fmt.Println("i::", i)
		// time.Sleep(time.Second)
	}
	fmt.Println("------------")
	for j := 0; j < 10; j++ {
		fmt.Println("j::", j)
		// time.Sleep(time.Second)
	}
	fmt.Println("------------")
	nomes := [3]string{"nicolas", "rocha", "Felippe"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, nome := range "nomes" {
		fmt.Println(indice, string(nome))
	}

	mapUser := map[string]string{
		"nome":      "Nicolas",
		"sobrenome": "Rocha",
	}

	for chave, valor := range mapUser {
		fmt.Println(chave, valor)
	}
}
