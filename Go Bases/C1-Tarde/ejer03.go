package main

/*
Ejercicio 3 - Préstamo

Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados
y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/
import "fmt"

var edad, sueldo, antiguedad int
var empleado bool

func main() {
	fmt.Println("Ejercicio 3")

	edad := 23
	antiguedad := 2
	empleado := true
	sueldo := 150000

	if edad > 22 && antiguedad > 1 && empleado == true {
		if sueldo >= 100000 {
			fmt.Println("Prestamo aprobado, No se le aplicaran intereses.")
		} else {
			fmt.Println("Prestamo aprobado, Se le aplicaran intereses de 15% anual.")
		}
	} else if edad <= 22 {
		fmt.Println("El usuario es menor de 22 años, no se puede efectuar el prestamo.")
	} else if antiguedad <= 1 {
		fmt.Println("El usuario tiene una antiguedad menor a 1 año, no se puede efectuar el prestamo.")
	} else if empleado == false {
		fmt.Println("El usuario no se encuentra empleado, no se puede efectuar el prestamo.")
	}

	//Resolucion con una funcion.
	fmt.Println(prestamo(25, 3, 99, true))

	//Resolucion con una funcion.
	fmt.Println(prestamo(18, 3, 99, true))

	//Resolucion con una funcion.
	fmt.Println(prestamo(25, 3, 101, false))
}

func prestamo(edad int, antiguedad int, sueldo int, empleado bool) string {

	if edad > 22 && antiguedad > 1 && empleado == true {

		if sueldo >= 100000 {
			return "Prestamo aprobado, No se le aplicaran intereses."
		} else {
			return "Prestamo aprobado, Se le aplicaran intereses de 15% anual."
		}

	} else if edad <= 22 {
		return "El usuario es menor de 22 años, no se puede efectuar el prestamo."

	} else if antiguedad <= 1 {
		return "El usuario tiene una antiguedad menor a 1 año, no se puede efectuar el prestamo."

	} else if empleado == false {
		return "El usuario no se encuentra empleado, no se puede efectuar el prestamo."

	} else {
		return "Hay un fallo en el ingreso de datos."
	}

}
