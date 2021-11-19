package main

/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para
administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
-Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de
la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia
		 en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento,
		 y un costo adicional  por envío de $2500.

Requerimientos:

-Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
-Crear una estructura “tienda” que guarde una lista de productos.

-Crear una interface “Producto” que tenga el método “CalcularCosto”
-Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.

-Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre
	y precio y devuelva un Producto.
-Se requiere una función “nuevaTienda” que devuelva un Ecommerce.

-Interface Producto:
-El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.

-Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos
 	y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

*/
import "fmt"

//-Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
type producto struct {
	tipo   string
	nombre string
	precio float64
}

//-Crear una estructura “tienda” que guarde una lista de productos.
type tienda struct {
	producto []producto
}

//-Crear una interface “Producto” que tenga el método “CalcularCosto”
type Producto interface {
	CalcularCosto() float64
}

//-Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
type Ecommerce interface {
	Total() float64
	Agregar(p producto) string
}

func main() {
	fmt.Println("Ejercicio 3")

	pro := producto{
		tipo:   "Pequeno",
		nombre: "Coca cola",
		precio: 200,
	}
	fmt.Println("\nUn producto: ", pro)

	//---------------------------------------------------------------------------------------

	//Creo un arreglo de productos
	var prto [3]producto

	prto[0].tipo = "Pequeña" // Asignacion campo a campo
	prto[0].nombre = "Polenta"
	prto[0].precio = 50

	prto[1].tipo = "Mediana" // Asignacion campo a campo
	prto[1].nombre = "Pack de azucar"
	prto[1].precio = 500

	prto[2] = producto{ // Asignacion mediante un objeto JSON
		tipo:   "Grande",
		nombre: "Pack de aceite",
		precio: 2000,
	}

	t := tienda{
		//producto: [producto{tipo:"Pequeno",nombre: "Coca cola",precio: 200}],
		producto: prto[:],
	}

	fmt.Println("\nTienda: ", t)

	//---------------------------------------------------------------------------------------

	t1 := nuevaTienda()
	fmt.Println("\nLa nueva tienda tiene: ", t1)
	fmt.Println("\nEl total de la tienda es: ", t1.Total())
	pNuevo := producto{
		tipo:   "Pequeno",
		nombre: "Coca cola",
		precio: 200,
	}
	fmt.Println("\nAgregar un producto: ", t1.Agregar(pNuevo))
	fmt.Println(t1)
	fmt.Println("\nEl total de la tienda es: ", t1.Total())
}

//-Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre
//	y precio y devuelva un Producto.
func nuevoProducto(tipo string, nombre string, precio float64) producto {

	p := producto{
		tipo:   tipo,
		nombre: nombre,
		precio: precio,
	}

	return p
}

/*-Se requiere una función “nuevaTienda” que devuelva un Ecommerce.*/
func nuevaTienda() Ecommerce {
	//Generar lista de productos y devolverlos en una tienda nueva

	var prto [2]producto
	prto[0] = producto{ // Asignacion mediante un objeto JSON
		tipo:   "Pequeña",
		nombre: "Aceitunas",
		precio: 100,
	}
	prto[1] = producto{ // Asignacion mediante un objeto JSON
		tipo:   "Pequeña",
		nombre: "Fideo codito",
		precio: 150,
	}
	t := tienda{producto: prto[:]}

	return t

}

//-Interface Producto:
//-El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
func (p producto) CalcularCosto() float64 {

	switch p.nombre {
	case "Pequeño":
		return p.precio

	case "Mediano":
		return p.precio + p.precio*0.03

	case "Grande":
		return p.precio + p.precio*0.06 + 2500

	default:
		return 0
	}
}

/*-Interface Ecommerce:
- El método “Total” debe retornar el precio total en base al costo total de los productos
	y los adicionales si los hubiera.
- El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda*/
func (t tienda) Total() float64 {

	var total float64
	for _, prod := range t.producto {
		total += prod.precio
	}

	return total
}

/**/
func (t tienda) Agregar(p producto) string {

	t.producto = append(t.producto, p)
	//msj := "Producto: ", p.nombre," agregado correctamente"
	return "Producto agregado correctamente"
}
