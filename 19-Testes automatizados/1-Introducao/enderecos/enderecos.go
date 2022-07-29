package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "estrada", "rodovia", "avenida"}

	letraMin := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(letraMin, " ")[0]
	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
			break
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tipo inválido"
}
