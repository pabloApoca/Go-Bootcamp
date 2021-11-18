package main

/*
Ejercicio 1 - Impuestos de salario

Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más
de $150.000 se le descontará además un 10%.
*/
import "fmt"

var sueldo, impuesto int

func main() {
	fmt.Println("Ejericio 1")
	calcularImpuesto(10000)
	calcularImpuesto(60500)
	calcularImpuesto(170700)
}

func calcularImpuesto(sueldo float64) {
	var impuesto float64
	if sueldo > 50000 && sueldo <= 149000 {
		impuesto = (sueldo / 100) * 17
		fmt.Println("Su impuesto es de: $", impuesto)
	} else if sueldo >= 150000 {
		impuesto = (sueldo / 100) * 27
		fmt.Println("Su impuesto es de: $", impuesto)
	} else {
		fmt.Println("Su impuesto es de: $ 0")
	}
}
