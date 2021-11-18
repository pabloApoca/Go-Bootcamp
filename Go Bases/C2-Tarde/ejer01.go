package main

/*
Ejercicio 1 - Registro de estudiantes

Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables
Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/
import "fmt"

type Estudiante struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func main() {
	fmt.Println("Ejercicio 1")

	estudiante := Estudiante{
		nombre:   "Pablo",
		apellido: "Lopez",
		dni:      11223344,
		fecha:    "18/04/1994",
	}

	estudiante.detalleEstudiante()

	estudiante2 := Estudiante{
		nombre:   "Juan",
		apellido: "Perez",
		dni:      22222222,
		fecha:    "01/01/1991",
	}
	estudiante2.detalleEstudiante()

}

func (a Estudiante) detalleEstudiante() {
	fmt.Println("\nEstudiante: \n Nombre: ", a.nombre, "\n Apellido: ", a.apellido, "\n Dni: ", a.dni, "\n Fecha: ", a.fecha)
}
