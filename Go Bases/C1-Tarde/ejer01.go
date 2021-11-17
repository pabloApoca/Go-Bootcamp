package main

//Ejercicio 1 - Letras de una palabra
//La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
//1-Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
//2-Luego imprimí cada una de las letras.
import "fmt"

var palabra string

func main() {
	fmt.Println("Ejercicio 1")
	palabra := "Arreglado"

	longuitud := len(palabra)
	fmt.Println("La palabra ", palabra, " tiene ", longuitud, " letras")

	//Opcion 1
	/*for i := 0; i < len(palabra); i++ {
		fmt.Println(string(palabra[i]))
	}*/

	//Opcion 2
	for indice, letra := range palabra {
		fmt.Printf("Posicion: %d - Letra: %c \n", indice, letra)
	}

}
