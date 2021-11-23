package main

/*
Ejercicio 4 - Ordenamiento
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
-	un arreglo de números enteros con 100 valores
-	un arreglo de números enteros con 1000 valores
-	un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand
package main

import (
   "math/rand"
   "fmt"
)

func main() {
   variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(10000)
}

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
)

func main() {
	fmt.Println("Ejercicio 4")
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	//variable3 := rand.Perm(10000)

	fmt.Println("\n\n", variable1)
	fmt.Println("\n\n", variable2)
	//fmt.Println("\n\n", variable3)

}

//-	Ordenamiento por inserción
func insercion(arrN []int) {

}

//-	Ordenamiento por burbuja
func burbuja(arrN []int) {

}

//-	Ordenamiento por selección
func seleccion(arrN []int) {

}