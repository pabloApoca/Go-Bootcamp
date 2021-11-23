package main

/*
Ejercicio 1 - Red social

Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando
información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el
mismo lugar en memoria para el main del programa y para las funciones:

La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña

Y deben implementarse las funciones:
-	cambiar nombre: me permite cambiar el nombre y apellido.
-	cambiar edad: me permite cambiar la edad.
-	cambiar correo: me permite cambiar el correo.
-	cambiar contraseña: me permite cambiar la contraseña.
*/
import (
	"fmt"
)

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func main() {
	fmt.Println("Ejercicio 1")
	fmt.Println("Agregar usuario: ")

	user := usuario{
		nombre:     "Pablo",
		apellido:   "Lopez",
		edad:       27,
		correo:     "pablo@gmail.com",
		contraseña: "123facil",
	}
	fmt.Println(user)

	var usPuntero *usuario

	usPuntero = &user

	fmt.Println("\nCambiar nombre:")
	fmt.Println(usPuntero.cambiarNombre("Martin", "Ledesma"))

	fmt.Println("Cambiar edad:")
	fmt.Println(usPuntero.cambiarEdad(30))

	fmt.Println("Cambiar correo:")
	fmt.Println(usPuntero.cambiarCorreo("pabloMaster@gmail.com"))

	fmt.Println("Cambiar contraseña:")
	fmt.Println(usPuntero.cambiarContraseña("456dificil"))

	fmt.Println("\nUser cambiado totalmente:")
	fmt.Println(user)

}

func (c *usuario) cambiarNombre(nombre, apellido string) string {
	c.nombre = nombre
	c.apellido = apellido
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarEdad(edad int) string {
	c.edad = edad
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarCorreo(correo string) string {
	c.correo = correo
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarContraseña(contraseña string) string {
	c.contraseña = contraseña
	return fmt.Sprintln(c)
}
