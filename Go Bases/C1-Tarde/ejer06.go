package main

/*
Ejercicio 6 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. }
Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
-Saber cuántos de sus empleados son mayores a 21 años.
-Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
-Eliminar a Pedro del mapa.
*/
import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
var mayores int

func main() {
	fmt.Println("Ejercicio 6")

	//Imprimir la edad de Benjamin.
	fmt.Println("\nLa edad de Benjamin es: ", employees["Benjamin"])

	//Cuántos de sus empleados son mayores a 21 años.
	for _, employee := range employees {
		if employee > 21 {
			mayores += 1
		}
	}
	fmt.Println("\nHay ", mayores, " empleados que son mayores a 21 años.")

	//Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	fmt.Println("\nAgrego a Federico de 25 años.")
	employees["Federico"] = 25
	fmt.Println(employees)

	//Eliminar a Pedro del mapa.
	fmt.Println("\nElimino a Pedro del Mapa.")
	delete(employees, "Pedro")
	fmt.Println(employees)

}
