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
import "fmt"

func main() {
	fmt.Println("Ejericio 4")
}
