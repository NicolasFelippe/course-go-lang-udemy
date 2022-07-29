package main

import "fmt"

func main() {
	fmt.Println("Maps")
	usuario := map[string]string{
		"nome":      "Nicolas",
		"sobrenome": "Rocha",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"user": {
			"nome":  "Nicolas",
			"idade": "2",
		},
	}
	delete(usuario2, "user")
	fmt.Println(usuario2)
}
