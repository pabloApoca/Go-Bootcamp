package main

/*
Ejercicio 3 - Calcular salario

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad
de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
por mes y la categoría, y que devuelva su salario.

*/
import (
	"errors"
	"fmt"
)

var categoria string
var salario int

func main() {
	fmt.Println("Ejericio 3")

	//Calcular salario 1
	salario, err := calcularSalario(120, "c")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario del empleado es: $", salario)
	}

	//Calcular salario 2
	salario2, err := calcularSalario(120, "b")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario del empleado es: $", salario2)
	}

	//Calcular salario 3
	salario3, err := calcularSalario(120, "A")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario del empleado es: $", salario3)
	}

	//Calcular salario 4
	salario4, err := calcularSalario(120, "z")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario del empleado es: $", salario4)
	}
}

func calcularSalario(minutos int, categoria string) (int, error) {
	var salario, salarioNeto int
	var horas int
	switch categoria {
	case "c", "C":
		horas = minutos / 60
		salarioNeto = horas * 1000

	case "b", "B":
		horas = minutos / 60
		salario = horas * 1500
		salarioNeto = salario + ((salario / 100) * 20)

	case "a", "A":
		horas = minutos / 60
		salario = horas * 3000
		salarioNeto = salario + ((salario / 100) * 50)
	default:
		return 0, errors.New("Categoria invalida.")

	}

	return salarioNeto, nil
}
