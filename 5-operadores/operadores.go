package main

import "fmt"

func main() {
	fmt.Println("operadores")
	/* ARITMETICOS
	+,-,*,/
	*/
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restodaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restodaDivisao)

	/*  ATRIBUIÇÃO */
	var var1 string = "var1"
	var2 := "var2"
	fmt.Println(var1, var2)

	/* OPERADORES RELACIONAIS	*/
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	/* OPERADORES LOGICOS */
	// and
	fmt.Println(false && true)
	// or
	fmt.Println(false || true)
	// negação
	fmt.Println(!false)

	/* OPERADORES UNÁRIOS */

	num1 := 10
	num1++
	fmt.Println(num1)
	num1--
	fmt.Println(num1)
	num1 += 10
	fmt.Println("soma", num1)
	num1 -= 10
	fmt.Println("subtração", num1)
	num1 *= 10
	fmt.Println("multiplicação", num1)
	num1 /= 2
	fmt.Println("divisao", num1)
	num1 %= 2
	fmt.Println("resto", num1)

	/* OPERADORES TERNARIO  não tem !!*/
	var text2 string
	if num1 >= 0 {
		text2 = "text"
	} else {
		text2 = "text-else"
	}

	fmt.Println(text2)

}
