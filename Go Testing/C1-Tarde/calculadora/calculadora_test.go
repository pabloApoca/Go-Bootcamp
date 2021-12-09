package calculadora

//Sobre el mismo file en el que venimos trabajando “calculadora_test.go” implementamos testify
//de la siguiente manera

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Sumar(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

//Las validaciones se hacen a través del package“assert”, que nos ofrece de forma sencilla efectuar las
//comparaciones y validaciones. Las funciones principales disponibles son:

// Validar Igualdad ->  assert.Equal(t,  123, 123, "deberían ser iguales")
// Validar desigualdad -> assert.NotEqual(t, 123, 456, "no deberían ser iguales")
// Validar Nulo Esperado (Bueno para errores) ->  assert.Nil(t, object)
//Validar No Nulo Esperado (Bueno para cuando esperamos algo) -> assert.NotNil(t, object)
