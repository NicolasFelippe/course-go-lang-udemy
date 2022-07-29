package main

import "fmt"

func main() {
	fmt.Println(somar(1, 1))

	var f = func(text string) string {
		fmt.Sprintln(text)
		return text
	}
	resultado := f("loucura")
	fmt.Println(resultado)
	cal1, cal2 := calculosMatematicos(30, 40)
	fmt.Println("calculosMatematicos:", cal1, cal2)

	// descarta o retorno _
	cal3, _ := calculosMatematicos(30, 40)
	fmt.Println("calculosMatematicos:", cal3)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Declara um tipo para dois parametros, sempre pega o ultimo tipo declarado.
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
