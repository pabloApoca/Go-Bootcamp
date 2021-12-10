package calculadora

//Mock
//El Mock, contrario a un Stub, no es aplicado para devolver valores exactos sino para comprobar
//todo el funcionamiento interno del método o código que se está probando. El Mock está más
//interesado en que métodos se han invocado, con que argumentos, cuando y con qué frecuencia.

//Un mock siempre es un espía y conoce lo que se se está testeando. Y las comprobaciones del test
//se aplican sobre el mock. Supongamos que ahora tenemos una función Sumar más compleja que
//llamaremos SumarRestricted(), en la que sólo se devolverá el resultado, si el proceso lo invoca
//un cliente autorizado para tal fin.

//zPor lo que recibe por parámetros, el nombre del cliente que está ejecutando el proceso, y la
//interfaz que valida si dicho cliente está autorizado.

//Para probar esta función necesitaremos un Mock que compruebe que la validación SumaEnabled()
//está siendo invocada correctamente y que además reciba el cliente correcto

type Config interface {
	SumaEnabled(cliente string) bool
}

// Función que recibe dos enteros y retorna la suma resultante
func SumarRestricted(num1, num2 int, config Config, cliente string) int {
	if !config.SumaEnabled(cliente) {
		return -99999
	}
	return num1 + num2

}
