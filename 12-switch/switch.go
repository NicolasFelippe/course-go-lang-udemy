package main

import "fmt"

func diaDaSemana(num int) string {

	var diaDaSemana string
	switch num {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough
	case 2:
		diaDaSemana = "Segunda-Feira"
	case 3:
		diaDaSemana = "Terça-feira"
	case 4:
		diaDaSemana = "Quarta-feira"
	case 5:
		diaDaSemana = "Quinta-feira"
	case 6:
		diaDaSemana = "Sexta-feira"
	case 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Numero inválido"
	}

	return diaDaSemana
}

func diaDaSemana2(num int) string {
	switch {
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda-Feira"
	case num == 3:
		return "Terça-feira"
	case num == 4:
		return "Quarta-feira"
	case num == 5:
		return "Quinta-feira"
	case num == 6:
		return "Sexta-feira"
	case num == 7:
		return "Sabado"
	default:
		return "Numero inválido"
	}
}

func main() {
	fmt.Println("Switch -------")
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana2(1))
}
