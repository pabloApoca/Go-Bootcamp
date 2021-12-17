package main

import (
	"fmt"
	"math"
)

type Persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Direccion string
}

type Circulo struct {
	Radio float64
	Area  float64
}

func main() {
	fmt.Println("Metodos")
	//p1 := Persona{"Pablo", "Lopez", 27, "Calle 3"}
	//fmt.Printf("\nP1: %+v, %T", p1, p1)
	//Concatenar(p1)
	//p1.MetodoConcatenar()

	c1 := Circulo{
		Radio: 3,
	}
	fmt.Printf("\n%+v\n", c1)
	c1.area()
	c1.perimetro()

	//areaFuncion(&c1)
	fmt.Printf("\n%+v", c1)
	c1.setRadio(5.7)
	fmt.Printf("\n%+v", c1)
}

func Concatenar(p Persona) {
	fmt.Println(p.Nombre, " ", p.Apellido)
}

func (p Persona) MetodoConcatenar() {
	fmt.Println(p.Nombre, " ", p.Apellido)
}

func (c *Circulo) area() {
	c.Area = math.Pi * c.Radio * c.Radio
	fmt.Println(c.Area)
}

func (c Circulo) perimetro() {
	fmt.Println(2 * math.Pi * c.Radio)
}

func areaFuncion(c *Circulo) {
	fmt.Println("\nDentro de areaFuncion")
	c.Area = math.Pi * c.Radio * c.Radio
	fmt.Println(c.Area)
}

func (c *Circulo) setRadio(nuevoRadio float64) {
	c.Radio = nuevoRadio
}
