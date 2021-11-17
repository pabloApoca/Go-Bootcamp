package main

/*
Ejercicio 2 - Descuento

Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos,
para ello necesitan una aplicación que les permita calcular el descuento en base a 2 variables,
su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento aplicado
y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.
*/
import "fmt"

var precio, descuento int

func main() {
	fmt.Println("Ejercicio 2")
	precio, descuento := 10000, 50
	fmt.Printf("El descuento de la prenda es: %d", hacerDescuento(precio, descuento))

}

func hacerDescuento(precio, descuento int) int {
	return (precio / 100) * descuento
}
