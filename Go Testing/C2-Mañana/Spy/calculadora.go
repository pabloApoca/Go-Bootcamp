package calculadora

//Spy
//En ocasiones es necesario comprobar o asegurarse de haber llamado a un método para dar el test como
//válido. Para esto utilizamos un Spy. Y la comprobación consiste en consultarle al Spy si el método
//en cuestión, fue invocado o utilizado durante la ejecución. De allí el nombre de este tipo de tests,
//es un espía que nos informa cuando algo sucede. Nuestro próximo test consiste en comprobar que
//efectivamente el método Log del objeto logger sea invocado durante la prueba.

type Logger interface {
	Log(string) error
}

// Función que recibe dos enteros, un objeto del tipo logger y retorna la   suma resultante
func Sumar(num1, num2 int, logger Logger) int {
	err := logger.Log("Ingreso a Función Sumar")
	if err != nil {
		return -99999
	}
	return num1 + num2
}
