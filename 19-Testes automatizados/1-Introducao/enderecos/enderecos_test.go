package enderecos

// Teste unidade

import (
	"testing"
)

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// roda os testes em paralelo
	// t.Parallel()
	cenariosDeTeste := []cenarioDeTest{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça ABC", "Tipo inválido"},
		{"", "Tipo inválido"},
		{"Estrada ABC", "Estrada"},
		{"RUA DOS BOBOS ", "Rua"},
		{"AVENIDA ABC", "Avenida"},
		{"RODOVIA ABC", "Rodovia"},
	}

	for _, cenario := range cenariosDeTeste {
		retorno := TipoDeEndereco(cenario.enderecoInserido)
		if retorno != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retorno,
				cenario.retornoEsperado)
		}
	}
}
