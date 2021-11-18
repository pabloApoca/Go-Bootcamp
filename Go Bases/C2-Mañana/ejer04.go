package main

/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de
los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una
cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

Ejemplo:

const (
   minimo = "minimo"
   promedio = "promedio"
   maximo = "maximo"
)

...

minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)

...

valorMinimo := minFunc(2,3,3,4,1,2,4,5)
valorPromedio := promFunc(2,3,3,4,1,2,4,5)
valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

*/
import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	fmt.Println("Ejericio 4")

	//Base calcular minimo
	valorMinimo := calcularMinimo(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("\nMinimo: ", valorMinimo)

	//Base calcular maximo
	valorMaximo := calcularMaximo(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("Maximo: ", valorMaximo)

	//Base calcular promedio
	valorPromedio := calcularPromedio(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("Promedio: ", valorPromedio)

	//---------------------------------------------------------------------------------------

	//Usamos la funcion operacion y le pasamos los datos de las constantes
	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion("maximo")
	//controlamos por si alguna sale con error
	if err != nil {
		fmt.Println(err)
	}

	//Llamos a cada metodo tomado como respuesta de la operacion para cada uno
	valorMinimo1 := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	valorPromedio1 := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo1 := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	//Imprimimos los valores
	fmt.Println("\nMinino de la operacion: ", valorMinimo1)
	fmt.Println("Maximo de la operacion: ", valorMaximo1)
	fmt.Println("Promedio de la operacion: ", valorPromedio1)

}

func calcularMinimo(numeros ...int) int {
	min := 10000
	for _, num := range numeros {
		if num < min {
			min = num
		}
	}
	return min
}

func calcularMaximo(numeros ...int) int {
	max := 0
	for _, num := range numeros {
		if num > max {
			max = num
		}
	}
	return max
}

func calcularPromedio(numeros ...int) int {
	var promedio int
	var total int
	for _, num := range numeros {
		total += num
	}
	promedio = (total / len(numeros))
	return promedio
}

func operacion(calculo string) (func(...int) int, error) {

	switch calculo {
	case "minimo":
		min := calcularMinimo
		return min, nil
	case "maximo":
		max := calcularMaximo
		return max, nil
	case "promedio":
		prom := calcularPromedio
		return prom, nil
	default:
		return nil, errors.New("No existe calculo con ese nombre.")
	}
}
