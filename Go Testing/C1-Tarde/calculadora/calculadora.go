package calculadora

//Para nuestro primer test, tomaremos el siguiente código como objeto de pruebas.

// Función que recibe dos enteros y retorna la suma resultante
func Sumar(num1, num2 int) int {
	//return num1 + num2

	// Esta función ahora devuelve un resultado INCORRECTO
	return num1 - num2

}

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Restar(num1, num2 int) int {
	return num1 - num2
}

//Finalmente procedemos a la ejecución del test:
//go test
