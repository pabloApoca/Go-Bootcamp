package main

/*
Ejercicio 1 - Impuestos de salario #1

1.	En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo “int”.
2.	Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary”
sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/
import (
	"fmt"
	"os"
)

type error interface {
	Error() string
}

type myCustomError struct {
	msg string
}

var salary, salary2 int

func main() {
	fmt.Println("Ejercicio 1")

	salary := 150000

	msg, err := salatyTest(salary) //llamamos a nuestra func
	if err != nil {                //hacemos una validación del valor de err
		fmt.Println(err) //si err no es nil, imprimimos el error y...
		os.Exit(1)       //utilizamos este método para salir del programa
	}
	fmt.Println("\n", msg)

	salary2 := 15000

	msg2, err := salatyTest(salary2) //llamamos a nuestra func
	if err != nil {                  //hacemos una validación del valor de err
		fmt.Println(err) //si err no es nil, imprimimos el error y...
		os.Exit(1)       //utilizamos este método para salir del programa
	}
	fmt.Println("\n", msg2)

}

func (e *myCustomError) Error() string {
	return fmt.Sprintf(e.msg)
}

func salatyTest(salary int) (string, error) {
	var msg string
	if salary < 150000 {
		return msg, &myCustomError{
			msg: "\nerror: el salario ingresado no alcanza el mínimo imponible.",
		}
	} else {
		return "Debe pagar impuesto", nil
	}
}
