package main

/*
Ejercicio 4 -  Impuestos de salario #4

Vamos a hacer que nuestro programa sea un poco más complejo y útil.

1.	Desarrolla las funciones necesarias para permitir a la empresa calcular:

a)	Salario mensual de un trabajador según la cantidad de horas trabajadas. En caso de que
el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto
de impuesto. La función que se ocupe de realizar este cálculo deberá retornar más de un valor,
incluyendo un error en caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o
un número negativo. El error deberá indicar “error: el trabajador no puede haber trabajado
menos de 80 hs mensuales”.

b)	Calcular el medio aguinaldo correspondiente al trabajador (fórmula de cálculo de aguinaldo:
	[mejor salario del semestre] dividido 12 y multiplicar el [resultado obtenido]
	por la [cantidad de meses trabajados en el semestre]). La función que realice el cálculo
	deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

2.	Desarrolla el código necesario para cumplir con las funcionalidades requeridas,
utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar
las validaciones de los retornos de error en tu función “main()”.
*/
import (
	"errors"
	"fmt"
)

type trabajador struct {
	horasTrabajadas int
	salario         float64
}

func main() {
	fmt.Println("Ejercicio 4")
	fmt.Println("\nCalcular salario de usuario y controlar horas trabajadas") //a
	t1 := trabajador{
		horasTrabajadas: 80,
		salario:         0,
	}
	salari, err := t1.calcularSalario(t1)
	if err != nil {
		fmt.Println(err)
	} else {
		t1.salario = salari
		fmt.Println("El salario total del trabajado es: ", salari)
		fmt.Println("Detalle: --> Horas trabajadas: ", t1.horasTrabajadas, " - Salario: ", t1.salario)
	}

	fmt.Println("\nCalcular aguinaldo de usuario y controlar meses ingresados") //b
	aguinaldo, err := t1.aguinaldont(t1, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El aguinaldo del usuario: ", t1, " es de $", aguinaldo)

	}

}

//1 a) Calcular salario mensual de un trabajador según la cantidad de horas trabajadas
func (t *trabajador) calcularSalario(tra trabajador) (float64, error) {
	var totalSalario float64
	tra.salario = float64(tra.horasTrabajadas) * 2000 // 80 * 2000 = 160000

	if tra.salario >= 150000 {
		totalSalario = (tra.salario - ((tra.salario / 100) * 10))
	} else {
		totalSalario = tra.salario
	}
	if tra.horasTrabajadas < 80 {
		//fmt.Println(errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales"))
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	return totalSalario, nil
}

//1 b) Calcular el medio aguinaldo correspondiente al trabajador
func (t *trabajador) aguinaldont(tr trabajador, mesesTrabajados int) (float64, error) {
	var aguinaldo float64
	aguinaldo = (tr.salario / 12) * float64(mesesTrabajados)
	if mesesTrabajados <= 0 {
		return 0, errors.New("error: ingrese un numero de meses mayor a 0")
	}

	return aguinaldo, nil
}
