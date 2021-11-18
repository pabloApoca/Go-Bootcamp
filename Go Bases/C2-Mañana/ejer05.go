package main

/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo
tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.

Se solicita:
--Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado
  y que retorne una función y un error (en caso que no exista el animal)
--Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

ejemplo:


const (
   perro = "perro"
   gato = "gato"
)

...

animalPerro, err := Animal(perro)
animalGato, err := Animal(gato)

...

var cantidad float64
cantidad += animalPerro(5)
cantidad += animalGato(8)


*/
import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamser    = "Hamster"
	tarantula = "TARANTULA"
)

func main() {
	fmt.Println("Ejericio 5")

	cPerro, err := animal(perro)
	cGato, err := animal(gato)
	cHamster, err := animal(hamser)
	cTarantula, err := animal(tarantula)
	//cPez, err := animal("pez")
	if err != nil {
		fmt.Println(err)
	}

	totalComidaPerro := cPerro(3)
	totalComidaGato := cGato(5)
	totalComidaHamster := cHamster(2)
	totalComidaTarantula := cTarantula(7)
	//tPez := cPez(3)

	fmt.Println("\nPara los perros se necesita: ", totalComidaPerro, " kilos de comida.")
	fmt.Println("\nPara los gatos se necesita: ", totalComidaGato, " kilos de comida.")
	fmt.Println("\nPara los hamsters se necesita: ", totalComidaHamster, " kilos/gramos de comida.")
	fmt.Println("\nPara las tarantulas se necesita: ", totalComidaTarantula, " kilos/gramos de comida.")
	//fmt.Println("\nPara las tarantulas se necesita: ", tPez, " kilos/gramos de comida.")

}

func calcularPerro(cantidad float64) float64 { //Pasado a float64
	var total float64
	total = cantidad * 10
	return total
}

func calcularGato(cantidad float64) float64 { //Pasado a float64
	var total float64
	total = cantidad * 5
	return total
}

func calcularHamster(cantidad float64) float64 {
	var total float64
	total = cantidad * 0.250
	return total
}

func calcularTarantura(cantidad float64) float64 {
	var total float64
	total = cantidad * 0.150
	return total
}

//Funcion principal
func animal(animal string) (func(float64) float64, error) {

	switch animal {
	case "PERRO", "perro", "Perro":
		calcPerro := calcularPerro
		return calcPerro, nil

	case "GATO", "gato", "Gato":
		calcGato := calcularGato
		return calcGato, nil

	case "HAMSTER", "hamster", "Hamster":
		calcHamster := calcularHamster
		return calcHamster, nil

	case "TARANTULA", "tarantura", "Tarantula":
		calcTarantura := calcularTarantura
		return calcTarantura, nil

	default:
		return nil, errors.New("No tenemos ese tipo de animal.")
	}
}
