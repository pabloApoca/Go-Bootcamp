package main

/*
Ejercicio 1 - Datos de clientes

Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones.
Para ello, cuentan con todo el detalle necesario en un archivo .txt.

1.	Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo,
no han pasado el archivo a leer por nuestro programa.

2.	Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt” (recuerda lo visto sobre el pkg “os”).
Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al
intentar leer un archivo que no existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Ejercicio 1")
	fmt.Println("\nejecución inicializada..")

	_, err := os.Open("customers.txt")
	if err != nil {

		fmt.Println("\nProbando dentro del if...")

		defer func() { //Utilizo defer para asegurarme que todas las instrucciones sean ejecutadas antes de la finalización de la ejecución de un programa.
			panic("el archivo indicado no fue encontrado o está dañado")
		}()

	}

	fmt.Println("\nejecución finalizada")

}

