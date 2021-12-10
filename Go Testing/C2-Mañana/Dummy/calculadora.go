package calculadora

//Dummy
//Es un objeto vacío que implementa una interfaz específica. Su uso, comportamiento o respuesta es
//irrelevante y no nos importa. Sólo lo usamos para satisfacer dependencias necesarias para la ejecución
//del código que estamos probando. Por ejemplo, supongamos que una función Sumar() requiere de un objeto
//logger para efectos de trazabilidad. La trazabilidad no nos importa, porque solo queremos probar el método.
//Pero necesitamos crear un logger dummy que realmente no haga nada, pero que nos permita ejecutar la prueba.

//Para probar esta función necesitaremos un Dummy del tipo logger, no nos importa que hace pero es
//requerido por la función.

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
