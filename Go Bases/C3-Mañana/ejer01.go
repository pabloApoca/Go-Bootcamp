package main

/*

Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
1- Implementar una funcionalidad para guardar un archivo de texto, con la información
	de productos comprados, separados por punto y coma (csv).
2- Debe tener el id del producto, precio y la cantidad.
3- Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type producto struct {
	id       int
	precio   float64
	cantidad int
}

func main() {
	fmt.Println("Ejercicio 1")

	var prto [3]producto

	prto[0].id = 1 // Asignacion campo a campo
	prto[0].precio = 200
	prto[0].cantidad = 8

	prto[1].id = 2 // Asignacion campo a campo
	prto[1].precio = 500
	prto[1].cantidad = 4

	prto[2] = producto{ // Asignacion mediante un objeto JSON
		id:       3,
		precio:   1000,
		cantidad: 9,
	}

	//Genero toda la lista de productos y la almaceno en un string para que sea leido por []byte
	res := fmt.Sprintln(prto)
	listaProductos := []byte(res)

	err := ioutil.WriteFile("./Productos.txt", listaProductos, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//La función ReadFile( filename string ) recibe como parámetro la dirección y nombre del archivo
	//en formato texto y nos devuelve el contenido del archivo en bytes o un error en caso que lo haya .
	db, err := os.ReadFile("./Productos.txt")
	fmt.Println("Lista de Productos:", string(db)) //Se tiene que encapsular en un string() para devolver el contenido del archivo

	//ReadAll(r Reader) lee desde r hasta un error o EOF y devuelve los datos que leyó y un error si lo hubiera.
	r := strings.NewReader("./Productos.txt")
	b, err := io.ReadAll(r)
	fmt.Println(string(b))

}
