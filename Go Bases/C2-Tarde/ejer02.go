package main

/*
Ejercicio 2 - Matrix

Una empresa de inteligencia artificial necesita tener una funcionalidad para crear
una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:

-Set:  Recibe una serie de valores de punto flotante e inicializa los valores
		en la estructura Matrix
-Print: Imprime por pantalla la matriz de una formas más visible
		(Con los saltos de línea entre filas)

La estructura Matrix debe contener los valores de la matriz, la dimensión del alto,
la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

*/
import "fmt"

type Matriz struct {
	valores []float64
	alto    int
	ancho   int
}

func main() {
	fmt.Println("Ejercicio 2")

	matrix := Matriz{
		valores: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		alto:    3,
		ancho:   3,
	}

	//hacer matrix.set(valores etc..)
	matrix.set(1, 2, 3, 4, 5) //ok
	//Matriz.set(matrix)
	fmt.Println("\nEl numero maximo es: ", Matriz.max(matrix))
	fmt.Println("\nEs cuadratica: ", Matriz.esCuadratica(matrix))
	fmt.Println()
	//matrix.set(1, 2, 3, 4, 5)
	Matriz.print(matrix)
}

func (m Matriz) set(valores ...float64) { //Ver con punteros
	m.valores = valores
}

func (m Matriz) esCuadratica() bool {
	if m.alto == m.ancho {
		return true
	} else {
		return false
	}
}

func (m Matriz) max() float64 {
	var max float64
	max = 1
	for _, valor := range m.valores {
		if valor > float64(max) {
			max = float64(valor)
		}
	}
	return max
}

func (m Matriz) print() {
	if len(m.valores) == 0 {
		fmt.Println("La matriz no tiene valores.")
	}
	for _, valor := range m.valores {

		fmt.Print(valor, "\t")
	}
}
