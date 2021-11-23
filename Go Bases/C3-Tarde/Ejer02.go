package main

/*
Ejercicio 2 - Ecommerce

Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main
del programa como en las funciones.

Se necesitan las estructuras:
-	Usuario: Nombre, Apellido, Correo, Productos (array de productos).
-	Producto: Nombre, precio, cantidad.

Se requieren las funciones:
-	Nuevo producto: recibe nombre y precio, y retorna un producto.
-	Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
-	Borrar productos: recibe un usuario, borra los productos del usuario.
*/

import (
	"fmt"
)

type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []producto
}

func main() {
	fmt.Println("Ejercicio 2")

	var listaDeProductos []producto
	p1 := producto{
		nombre:   "Memoria ram",
		precio:   14500,
		cantidad: 2,
	}
	p2 := producto{
		nombre:   "Disco ssd",
		precio:   12500,
		cantidad: 1,
	}
	p3 := producto{
		nombre:   "Disco rigido",
		precio:   7000,
		cantidad: 1,
	}

	listaDeProductos = append(listaDeProductos, p1, p2, p3)

	user1 := usuario{
		nombre:    "Pablo",
		apellido:  "Lopez",
		correo:    "Pablo@gmail.com",
		productos: listaDeProductos,
	}

	fmt.Println("Usuario numero 1:")
	fmt.Println(user1)
	//-------------------------------------

	fmt.Println("\nNuevo producto:")
	newProd := nuevoProducto("Fuente 700w", 27000)
	fmt.Println(newProd)
	//-------------------------------------

	fmt.Println("\nAgrego un producto a un usuario con la cantidad que indique:")
	var punteroUser *usuario
	punteroUser = &user1
	var punteroProd *producto
	punteroProd = &newProd

	agregarProducto(punteroUser, punteroProd, 3)
	fmt.Println(user1)
	//-------------------------------------

	fmt.Println("\nElimino la lista de productos del usuario:")
	borrarProducto(&user1)
	fmt.Println(user1)

}

//-	Nuevo producto: recibe nombre y precio, y retorna un producto.
func nuevoProducto(nombre string, precio float64) producto {
	p := producto{
		nombre:   nombre,
		precio:   precio,
		cantidad: 1,
	}
	return p
}

//-	Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
func agregarProducto(user *usuario, prod *producto, cant int) {
	prod.cantidad = cant
	user.productos = append(user.productos, *prod)
}

//-	Borrar productos: recibe un usuario, borra los productos del usuario.
func borrarProducto(user *usuario) {
	user.productos = append(user.productos[:0], user.productos[len(user.productos):]...)
}
