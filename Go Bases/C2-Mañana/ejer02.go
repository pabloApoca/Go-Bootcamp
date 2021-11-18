package main

/*
Ejercicio 2 - Calcular promedio

Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva
el promedio y un error en caso que uno de los números ingresados sea negativo
*/
import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Ejericio 2")

	//Calcular promedio 1
	promedio, erro := calcularPromedio(3, 3, 3, -3, 3)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("El promedio del alumno es:", promedio)
	}

	//Calcular promedio 2
	promedio2, erro := calcularPromedio(7, 8, 9, 10, 6)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("El promedio del alumno es:", promedio2)
	}

	//Calcular promedio 3
	promedio3, erro := calcularPromedio()
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("El promedio del alumno es:", promedio3)
	}
}

func calcularPromedio(notas ...int) (int, error) {
	var totalDeNotas int

	if len(notas) == 0 {
		return 0, errors.New("No se puede dividir por 0.")
	}

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("No se admiten numeros negativos.")
		}

		totalDeNotas += nota
	}

	return totalDeNotas / len(notas), nil
}
