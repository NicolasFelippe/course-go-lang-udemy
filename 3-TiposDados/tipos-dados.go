package main

import (
	"errors"
	"fmt"
)

// Inteiros
/*
int - utiliza arquitetura do PC conforme os bits 64 ou 32
int8
int16
int32
int64

uint - unsigned não aceita numero com sinal
uint8,16,32,64

- alias
	rune  = int32 //
	byte = uint8
*/

/*
	** FLOATS **
	float32
	float64

*/

/* ** STRINGS **
string
não tem char
*/
func main() {
	/* INTEGER */
	var numero int16 = 10000
	fmt.Println(numero)

	/* FLOAT */
	var float float32 = 103.10
	fmt.Println(float)

	float2 := 1303.10
	fmt.Println(float2)

	/* String */
	var str1 string = "teste"
	fmt.Println(str1)

	str2 := "teste"
	fmt.Println(str2)

	/* CHAR */

	char := 'B' // retorna numero da tabela ASC
	fmt.Println(char)

	/*  valor inicial vars 0, "", nil */

	var txt string
	var num11 int
	fmt.Println(txt, num11)

	/* Booleans */
	var booleanof bool // = false
	var booleano bool = false
	fmt.Println(booleano, booleanof)

	/* Error */

	var erro error // = nil
	var erro1 error = errors.New("Erro interno!")
	fmt.Println(erro, erro1)
}
