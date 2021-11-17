package main

/*
Ejercicio 4 - A qué mes corresponde

Realizar una aplicación que contenga una variable con el número del mes.
1- Según el número, imprimir el mes que corresponda en texto.
2- ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
Ej: 7, Julio
*/
import "fmt"

var mes int
var meses = map[int]string{}

func main() {
	fmt.Println("Ejercicio 4")

	meses[1] = "Enero"
	meses[2] = "Febrero"
	meses[3] = "Marzo"
	meses[4] = "Abril"
	meses[5] = "Mayo"
	meses[6] = "Junio"
	meses[7] = "Julio"
	meses[8] = "Agosto"
	meses[9] = "Septiembre"
	meses[10] = "Octubre"
	meses[11] = "Noviembre"
	meses[12] = "Diciembre"

	//fmt.Println(len(meses)) //Para saber la cantidad de meses
	//fmt.Println(meses)
	for indice, mes := range meses {
		fmt.Println("Clave: ", indice, " Valor: ", mes)
	}

	//Resolucion 1
	mes := 7
	fmt.Println("El ", mes, " es el mes de ", meses[mes])

	//Resolucion 2
	fmt.Println(calcularMes(4))

}

func calcularMes(mes int) string {

	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	elMes := meses[mes]
	mensaje := "El mes de: " + elMes
	return mensaje
}
