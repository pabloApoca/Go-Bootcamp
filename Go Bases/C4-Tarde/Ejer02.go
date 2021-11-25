package main

/*
Ejercicio 2 - Registrando clientes

El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
Los datos requeridos para registrar a un cliente son:
-	Legajo
-	Nombre y Apellido
-	DNI
-	Número de teléfono
-	Domicilio

●	Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los
restantes gastos. Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como
valor a “Legajo”. Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa
la ejecución y aborte.

●	Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe. Para ello, necesitas leer
los datos de un archivo .txt. En algún lugar de tu código, implementa la función para leer un archivo llamado
“customers.txt” (como en el ejercicio anterior, este archivo no existe, por lo que la función que intente
	leerlo devolverá un error). Debes manipular adecuadamente ese error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la
ejecución del programa normalmente.

●	Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar
que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al
menos, dos valores. Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por
parámetro algún valor cero (recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).

●	Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes
mensajes: “Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos
abiertos” (en ese orden). Utiliza defer para cumplir con este requerimiento.

Requerimientos generales:
●	Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
●	Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error (por ejemplo las
	que intenten leer archivos).
Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello
(realiza también la validación pertinente para el caso de error retornado).

*/
import (
	"errors"
	"fmt"
	"os"
)

/*
Los datos requeridos para registrar a un cliente son:
-	Legajo
-	Nombre y Apellido
-	DNI
-	Número de teléfono
-	Domicilio*/
type cliente struct {
	legajo          int
	nombreYApellido string
	dni             int
	telefono        int
	domicilio       string
}

var listaDeClientes []cliente

func main() {
	fmt.Println("Ejercicio 2")

	c1 := cliente{
		legajo:          0,
		nombreYApellido: "Pablo Lopez",
		dni:             11223344,
		telefono:        12345678,
		domicilio:       "Pasaje 3",
	}

	c1.generarLegajo(&listaDeClientes) //Genero el legajo en base a la lista de clientes y se le asigno al mismo cliente
	//●	Tarea 2 - Hago como que leo un archivo para ver si ya existe el cliente
	_, err := os.Open("customers.txt")
	if err != nil {
		defer func() { //Utilizo defer para asegurarme que todas las instrucciones sean ejecutadas antes de la finalización de la ejecución de un programa.
			panic("error: el archivo indicado no fue encontrado o está dañado")
		}()

	} //Termino tarea 2
	listaDeClientes = append(listaDeClientes, c1)
	fmt.Println(c1)

	c2 := cliente{
		legajo:          0,
		nombreYApellido: "Diego Maradona",
		dni:             111111111,
		telefono:        1010101010,
		domicilio:       "El paraiso",
	}

	c2.generarLegajo(&listaDeClientes) //Genero el legajo en base a la lista de clientes y se le asigno al mismo cliente
	fmt.Println(c2)
	listaDeClientes = append(listaDeClientes, c2)

	//Funcio para tarea 3
	c, err := registrarCliente(0, "Juan Perez", 8888888, 42770011, "Pasaje 2")
	if err != nil {
		defer func() { //Utilizo defer para asegurarme que todas las instrucciones sean ejecutadas antes de la finalización de la ejecución de un programa.
			fmt.Println(err)
		}()
	}
	fmt.Println(c)

	fmt.Println("\n\tFin de la ejecución")
	fmt.Println("\tSe detectaron varios errores en tiempo de ejecución")
	fmt.Println("\tNo han quedado archivos abiertos")
	fmt.Println()
}

//●	Tarea 1 - Genero el legajo en base a una lista de clientes
func (c *cliente) generarLegajo(listaDeClientes *[]cliente) {

	var id int
	id = len(*listaDeClientes) + 1
	c.legajo = id

}

//●	Tarea 3 - Registro el cliente y verifico que todo los campos sean correctos
func registrarCliente(legajo int, nombreYApellido string, dni int, telefono int, domicilio string) (cliente, error) {

	c := cliente{
		legajo:          legajo,
		nombreYApellido: nombreYApellido,
		dni:             dni,
		telefono:        telefono,
		domicilio:       domicilio,
	}

	if legajo <= 0 {
		return c, errors.New("error: ingrese un numero de meses mayor a 0")
	}
	if nombreYApellido == "0" {
		return c, errors.New("error: ingrese nombre y apellido")
	}
	if dni == 0 {
		return c, errors.New("error: ingrese un numero de dni")
	}
	if telefono == 0 {
		return c, errors.New("error: ingrese un numero de telefono")
	}
	if domicilio == "0" {
		return c, errors.New("error: ingrese un domicilio")
	}

	return c, nil
}
