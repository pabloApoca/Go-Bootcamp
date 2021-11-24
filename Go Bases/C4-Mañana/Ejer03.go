package main

/*
Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje
de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el
	salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/
import (
	"fmt"
)

var salary, salary2 int //Domain Driven Design

func main() {
	fmt.Println("Ejercicio 3")

	salary := 150001
	if salary < 150000 {
		err := fmt.Errorf("\nerror: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(err)
		return
	}
	fmt.Println("\nDebe pagar impuesto")

	salary2 := 149000
	if salary2 < 150000 {
		err := fmt.Errorf("\nerror: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary2)
		fmt.Println(err)
		return
	}
	fmt.Println("nDebe pagar impuesto")
}
