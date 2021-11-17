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
	fmt.Printf("La prenda con decuento del %d sale: %d", descuento, hacerDescuento(precio, descuento))

	precio2, descuento2 := 7000, 25
	fmt.Printf("\nLa prenda con decuento del %d sale: %d", descuento, hacerDescuento(precio2, descuento2))

}

func hacerDescuento(precio, descuento int) int {
	descuentoAPrenda := (precio / 100) * descuento
	return precio - descuentoAPrenda
}
