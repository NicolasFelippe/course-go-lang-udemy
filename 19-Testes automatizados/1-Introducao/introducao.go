package main

import (
	"fmt"
	"intruducao-testes/enderecos"
)

// executa em todos modulos
// go test ./...
// go test -v
// go test --cover
// go test --coverprofile cobertura.txt
// go tool cover --func=cobertura.txt
// go tool cover --html =cobertura.txt

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
