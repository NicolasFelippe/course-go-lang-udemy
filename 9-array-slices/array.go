package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays e slices")

	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [2]string{"teste1", "teste2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	/*  SLICE */

	slice1 := []int{1, 2}
	// slice1[2] = 12
	slice1 = append(slice1, 333)
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))

	slice2 := array1[2:4]
	fmt.Println("slice2", slice2)
	array1[3] = 333
	fmt.Println("array1", array1)
	fmt.Println(slice2)

	/* ARRAYS INTERNOS */
	fmt.Println("---------ARRAYS INTERNOS-------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 55)
	slice3 = append(slice3, 56)
	fmt.Println(slice3)

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
