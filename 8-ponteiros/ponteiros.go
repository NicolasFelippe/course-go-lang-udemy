package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)
	var1++
	fmt.Println(var1, var2)

	var var3 int
	var ponteiro1 *int
	fmt.Println(var3, ponteiro1)

	/*
		&variavel = pega o valor de memoria
		&variavel = lê o endereço de memoria
	*/
	var3 = 100
	ponteiro1 = &var3
	fmt.Println(var3, ponteiro1)

	var3 = 150
	fmt.Println(var3, *ponteiro1)
}
