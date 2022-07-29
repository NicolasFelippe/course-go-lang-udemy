package main

import "fmt"

func main() {
	var var1 string = "alguma coisa"
	var2 := "alguma coisa 2" // inferencia de tipo

	fmt.Println(var1)
	fmt.Println(var2)

	// declarar varias vars
	var (
		var3 string = "var3"
		var4 string = "var4"
	)
	fmt.Println(var3)
	fmt.Println(var4)

	var5, var6 := "var5", "var6"

	fmt.Println(var5)
	fmt.Println(var6)

	// const vars

	const const1 string = "const1"
	fmt.Println(const1)

	// troca valores sem var aux.
	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}
