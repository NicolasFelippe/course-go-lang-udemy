package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")

	func() {
		fmt.Println("anonima")
	}()

	func(txt string) {
		fmt.Println(txt)
	}("anonima 222")

	retorno := func(txt string) string {
		return fmt.Sprintf("Recebido -> %s", txt)
	}("anonima 222")

	fmt.Println(retorno)
}
