package main

/*
Ejercicio 5 - Listado de nombres

A) Una profesora de la universidad quiere tener un listado con todos sus estudiantes.
Es necesario crear una aplicación que contenga dicha lista.
Estudiantes:
Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.

B) Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar
el código que escribiste inicialmente.
Estudiante:
Gabriela
*/
import "fmt"

var listEstudents []string

func main() {
	fmt.Println("Ejercicio 5")

	listEstudents := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	//Muestro la lista de Estudiantes
	fmt.Println(listEstudents)

	//Agrego la nueva estudiante y muestro
	listEstudents = append(listEstudents, "Gabriela")
	fmt.Println(listEstudents)

	//Agrego un nuevo estudiante y muestro
	listEstudents = append(listEstudents, "Pablo")
	fmt.Println(listEstudents)

	fmt.Println("En la clase hay: ", len(listEstudents), " Estudiantes.")
}
