package main

/*
Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
en reemplazo de “Error()”,  se implemente “errors.New()”.

*/
import (
	"errors"
	"fmt"
)

var salary, salary2 int

func main() {
	fmt.Println("Ejercicio 2")

	salary := 150001
	if salary < 150000 {
		fmt.Println(errors.New("\nerror: el salario ingresado no alcanza el mínimo imponible."))
		return
	}
	fmt.Println("\nDebe pagar impuesto")

	salary2 := 149000
	if salary2 < 150000 {
		fmt.Println(errors.New("\nerror: el salario ingresado no alcanza el mínimo imponible."))
		return
	}
	fmt.Println("nDebe pagar impuesto")

}
