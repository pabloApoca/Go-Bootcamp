package main

import (
	"fmt"
)

type Persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Direccion string
}

func main() {
	fmt.Println("Estructuras")

	//Primera Forma
	p1 := Persona{"Pablo", "Lopez", 27, "Calle 3"}
	fmt.Printf("\nP1: %+v, %T", p1, p1)

	//Segunda Forma
	p2 := Persona{}
	p2.Nombre = "Juan"
	fmt.Printf("\n\nP2: %+v, %T", p2, p2)

	//Tercera Forma
	p3 := Persona{
		Nombre:    "Mariano",
		Direccion: "Calle 1",
		Apellido:  "Bustamante",
	}
	fmt.Printf("\n\nP3: %+v, %T", p3, p3)
	fmt.Printf("\nEl nombre de p3 es: %+v", p3.Nombre)

	var p4 Persona
	fmt.Printf("\n\nP4: %+v, %T", p4, p4)
	fmt.Printf("\nP4: El largo del nombre: %d", len(p4.Nombre))
	p4.Nombre = "Esteban"
	fmt.Printf("\nP4: El largo del nombre: %d", len(p4.Nombre))

}
