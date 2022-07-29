package main

import (
	"fmt"
	"time"
)

func main() {

	//  COCORRÊNCIA != PARALELISMO
	go escrever("Olá mundo")
	escrever("Programando em go")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
