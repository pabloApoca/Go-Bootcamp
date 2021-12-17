package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Tipo de dato")

	var numero1 int
	type entero int
	numero1 = 25
	var numero2 entero
	numero2 = 32

	fmt.Printf("\nNumero 1: %v, %T, %d", numero1, numero1, unsafe.Sizeof(numero1))
	fmt.Printf("\nNumero 1: %v, %T, %d", numero2, numero2, unsafe.Sizeof(numero2))

}
