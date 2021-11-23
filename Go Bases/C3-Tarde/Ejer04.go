package main

/*
Ejercicio 4 - Ordenamiento
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
-	un arreglo de números enteros con 100 valores
-	un arreglo de números enteros con 1000 valores
-	un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand

Se debe realizar el ordenamiento de cada una por:
-	Ordenamiento por inserción
-	Ordenamiento por burbuja
-	Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento
fue mejor para cada arreglo


*/
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Ejercicio 4")
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	//fmt.Println("\n\n", variable1)
	//fmt.Println("\n\n", variable2)
	//fmt.Println("\n\n", variable3)
	//fmt.Println("\n\n", seleccion(variable1))

	//Creamos 3 canales para cada go routine de cada ordenamiento y los devolvemos
	cInser := make(chan time.Duration)
	go ordenarArreglosInsercion(variable1, variable2, variable3, cInser)
	totalInser := <-cInser
	fmt.Println("\nOrdenar arreglos por Insercion dura: ", totalInser)

	cBur := make(chan time.Duration)
	go ordenarArreglosBurbuja(variable1, variable2, variable3, cBur)
	totalBur := <-cBur
	fmt.Println("\nOrdenar arreglos por Burbuja dura: ", totalBur)

	cSelec := make(chan time.Duration)
	go ordenarArreglosSeleccion(variable1, variable2, variable3, cSelec)
	totalSelec := <-cSelec
	fmt.Println("\nOrdenar arreglos por Burbuja dura: ", totalSelec)

}

//-	Ordenamiento por inserción
func insercion(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
		auxiliar = ListaDesordenada[i]
		for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
			ListaDesordenada[j+1] = ListaDesordenada[j]
			ListaDesordenada[j] = auxiliar
		}
	}
	return ListaDesordenada
}

//-	Ordenamiento por burbuja
func burbuja(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		for j := 0; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] > ListaDesordenada[j] {
				auxiliar = ListaDesordenada[i]
				ListaDesordenada[i] = ListaDesordenada[j]
				ListaDesordenada[j] = auxiliar
			}
		}
	}
	return ListaDesordenada
}

//-	Ordenamiento por selección
func seleccion(ListaDesordenada []int) []int {
	var auxiliar int

	for i := 0; i < len(ListaDesordenada); i++ {
		for j := i; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] > ListaDesordenada[j] {
				auxiliar = ListaDesordenada[j]
				ListaDesordenada[j] = ListaDesordenada[i]
				ListaDesordenada[i] = auxiliar
			}
		}
	}
	return ListaDesordenada
}

func ordenarArreglosInsercion(arreglo1 []int, arreglo2 []int, arreglo3 []int, c chan time.Duration) time.Duration {
	var startTime = time.Now()
	insercion(arreglo1)
	insercion(arreglo2)
	insercion(arreglo3)
	var duration = time.Since(startTime) // or: time.Now() - startTime
	c <- duration

	return duration
}

func ordenarArreglosBurbuja(arreglo1 []int, arreglo2 []int, arreglo3 []int, c chan time.Duration) time.Duration {
	var startTime = time.Now()
	burbuja(arreglo1)
	burbuja(arreglo2)
	burbuja(arreglo3)
	var duration = time.Since(startTime) // or: time.Now() - startTime
	c <- duration

	return duration
}

func ordenarArreglosSeleccion(arreglo1 []int, arreglo2 []int, arreglo3 []int, c chan time.Duration) time.Duration {
	var startTime = time.Now()
	seleccion(arreglo1)
	seleccion(arreglo2)
	seleccion(arreglo3)
	var duration = time.Since(startTime) // or: time.Now() - startTime
	c <- duration

	return duration
}
