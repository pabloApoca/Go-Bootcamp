package main

/*
Ejercicio 4 -  Impuestos de salario #4

Vamos a hacer que nuestro programa sea un poco más complejo y útil.

1.	Desarrolla las funciones necesarias para permitir a la empresa calcular:

a)	Salario mensual de un trabajador según la cantidad de horas trabajadas. En caso de que
el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto
de impuesto. La función que se ocupe de realizar este cálculo deberá retornar más de un valor,
incluyendo un error en caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o
un número negativo. El error deberá indicar “error: el trabajador no puede haber trabajado
menos de 80 hs mensuales”.
b)	Calcular el medio aguinaldo correspondiente al trabajador (fórmula de cálculo de aguinaldo:
	[mejor salario del semestre] dividido 12 y multiplicar el [resultado obtenido]
	por la [cantidad de meses trabajados en el semestre]). La función que realice el cálculo
	deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

2.	Desarrolla el código necesario para cumplir con las funcionalidades requeridas,
utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar
las validaciones de los retornos de error en tu función “main()”.
*/
import (
	"fmt"
)

func main() {
	fmt.Println("Ejercicio 4")

}
