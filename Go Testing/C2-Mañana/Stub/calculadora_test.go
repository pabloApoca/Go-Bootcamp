package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

// se crea un un struct stubLogger
type stubLogger struct{}

//  Se escribe las funciones necesarias para que stubLogger retorne exactamente lo que necesitamos
func (s *stubLogger) Log(string) error {
	return errors.New("error desde stub")
}

func TestSumarError(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := -99999
	// Se genera el objeto stub a usar para satisfacer la necesidad de la función Sumar
	myStub := &stubLogger{}
	// Se ejecuta el test
	resultado := Sumar(num1, num2, myStub)
	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
