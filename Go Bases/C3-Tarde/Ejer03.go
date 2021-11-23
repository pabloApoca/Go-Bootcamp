package main

/*
Ejercicio 3 - Calcular Precio

Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren
que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
-	Productos: nombre, precio, cantidad.
-	Servicios: nombre, precio, minutos trabajados.
-	Mantenimiento: nombre, precio.

Se requieren 3 funciones:
-	Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
-	Sumar Servicios: recibe un array de servicio y devuelve el precio total
(precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
-	Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final
(sumando el total de los 3).

*/
import (
	"fmt"
)

/*Se requieren 3 estructuras:
-	Productos: nombre, precio, cantidad.
-	Servicios: nombre, precio, minutos trabajados.
-	Mantenimiento: nombre, precio.*/

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicios struct {
	nombre        string
	precio        float64
	minTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func main() {
	fmt.Println("Ejercicio 3")

	var listaDeProductos []Productos //Listado de Productos
	prod1 := Productos{
		nombre:   "Coca cola",
		precio:   200,
		cantidad: 4,
	}
	prod2 := Productos{
		nombre:   "Fernet",
		precio:   800,
		cantidad: 2,
	}
	prod3 := Productos{
		nombre:   "Hielo",
		precio:   300,
		cantidad: 4,
	}
	listaDeProductos = append(listaDeProductos, prod1, prod2, prod3)

	var listaDeServicios []Servicios //Listado de Servicios
	serv1 := Servicios{
		nombre:        "Masajes",
		precio:        2000,
		minTrabajados: 180,
	}
	serv2 := Servicios{
		nombre:        "Peluqueria",
		precio:        500,
		minTrabajados: 120,
	}
	serv3 := Servicios{
		nombre:        "Maquillaje",
		precio:        700,
		minTrabajados: 240,
	}
	listaDeServicios = append(listaDeServicios, serv1, serv2, serv3)

	var listaDeMantenimiento []Mantenimiento //Listado de Mantenimiento
	mant1 := Mantenimiento{
		nombre: "Limpieza",
		precio: 700,
	}
	mant2 := Mantenimiento{
		nombre: "Lavado de auto",
		precio: 1000,
	}
	mant3 := Mantenimiento{
		nombre: "Gimnasios",
		precio: 300,
	}
	listaDeMantenimiento = append(listaDeMantenimiento, mant1, mant2, mant3)

	//Imprimir listas y calcular totales
	fmt.Println("\nLista de Productos: ", listaDeProductos) //----Productos
	c1Pro := make(chan float64)                             //Creo un canal
	go sumarProductos(listaDeProductos, c1Pro)              //Mando la lista con el canal
	totaP := <-c1Pro                                        // recibimos el valor del canal
	fmt.Println("Total: ", totaP)

	fmt.Println("\nLista de Servicios: ", listaDeServicios) //----Servicios
	c2Ser := make(chan float64)
	go sumarServicios(listaDeServicios, c2Ser)
	totalS := <-c2Ser
	fmt.Println("Total: ", totalS)

	fmt.Println("\nLista de Mantenimiento: ", listaDeMantenimiento) //----Mantenimiento
	c3Mant := make(chan float64)
	go sumarMantenimiento(listaDeMantenimiento, c3Mant)
	totalM := <-c3Mant
	fmt.Println("Total: ", totalM)

	fmt.Println("\nEl total de todo es: ", totaP+totalS+totalS)

}

//-	Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
func sumarProductos(productos []Productos, c chan float64) {
	var total float64

	for _, producto := range productos {
		total += (producto.precio * float64(producto.cantidad))
	}
	c <- total
}

//-	Sumar Servicios: recibe un array de servicio y devuelve el precio total
//(precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
func sumarServicios(servicios []Servicios, s chan float64) {

	var total float64

	for _, servicio := range servicios {
		horas := servicio.minTrabajados / 30
		total += (servicio.precio * float64(horas))
	}
	s <- total
}

//-	Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.*/

func sumarMantenimiento(mantenimientos []Mantenimiento, m chan float64) {
	var total float64

	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.precio
	}

	m <- total
}

/*func sumTotal(lprod []Productos, lserv []Servicios, lmant []Mantenimiento) float64 {




	return go sumarProductos(lprod) + go sumarServicios(lserv) + go sumarMantenimiento(lmant)
}*/
