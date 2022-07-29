package main

import (
	"fmt"
	"os"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("------", pwd, "-------")
	generica("teste")
}
