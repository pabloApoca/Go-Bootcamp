package calculadora

//Stub
//El propósito de un Stub es el de proveer valores concretos para guiar al test en una determinada dirección.
//Implementa métodos y devuelve valores requeridos para el test. Por ejemplo:  en la función “Sumar” -
//descrita previamente - existe una condición, en la que si el Logger.Log, retorna un error, la suma no se
//ejecuta sino que retorna el valor -99999.

//Para probar esta condición no es suficiente con crear un Dummy, sino que es necesario que el objeto que
//estamos simulando devuelve específicamente un error. Esto es precisamente lo que hace el Stub.

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
