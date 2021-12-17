package main

import (
	"encoding/json"
	"fmt"
)

type Fecha struct {
	Dia, Mes, Anio int
}

type Persona struct {
	Nombre    string `json:"user_nombre"`
	Apellido  string `json:"user_apellido"`
	Edad      int    `json:"edad"`
	Direccion string
	FechaNac  Fecha
}

func main() {
	fmt.Println("Estructuras 2")

	//Primera Forma
	p1 := Persona{"Pablo", "Lopez", 27, "Calle 3", Fecha{18, 4, 1994}}
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

	miJson, _ := json.Marshal(p3)
	fmt.Println()
	fmt.Println(string(miJson))

	/*var p4 Persona
	fmt.Printf("\n\nP4: %+v, %T", p4, p4)
	fmt.Printf("\nP4: El largo del nombre: %d", len(p4.Nombre))
	p4.Nombre = "Esteban"
	fmt.Printf("\nP4: El largo del nombre: %d", len(p4.Nombre))

	p5 := Persona{
		Nombre:    "Mariano",
		Direccion: "Calle 1",
		Apellido:  "Bustamante",
		FechaNac: Fecha{
			Dia:  2,
			Mes:  5,
			Anio: 1995,
		},
	}
	fmt.Printf("\n\nP5: %+v, %T", p5, p5)
	p5.FechaNac.Dia = 12
	fmt.Printf("\n\nP5: %+v, %T", p5, p5)

	fmt.Printf("\nReflect\n")
	tipo := reflect.TypeOf(p5)
	valor := reflect.ValueOf(p5)
	/*fmt.Println(tipo.Field(0))
	fmt.Println(tipo.Field(1))
	fmt.Println(tipo.Field(2))
	fmt.Println(tipo.NumField())*/

	/*for i := 0; i < tipo.NumField(); i++ {
		fmt.Printf("\nMiembro: %+v\n", tipo.Field(i))
		fmt.Printf("\nValor: %+v\n", valor.Field(i))

	}*/

}
