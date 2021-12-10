package calculadora

//Fake
//Los Fake son objetos que contienen cierta lógica de negocio adentro. Es una especie de “simulador”
//que debe generar respuestas distintas de acuerdo a distintos escenarios. Esto permite comprobar
//validaciones o comportamientos asociados al negocio, en los que además no podemos usar datos o
//escenarios productivos. Un Fake se distingue del resto de los test Double, ya que ningún otro
//contiene lógica de negocio. Son tests que tienden a crecer en complejidad en la medida que más
//lógica contengan.

//Tomando como objeto de pruebas nuevamente la función SumarRestricted, vamos a generar un Fake
//de la interfaz Config, cuyo método SumaEnabled devuelve “true” para un cliente específico. En
//otras palabras, queremos que solo un cliente pueda acceder al método Suma, el resto, debe recibir
//como resultado el  número -99999.

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
