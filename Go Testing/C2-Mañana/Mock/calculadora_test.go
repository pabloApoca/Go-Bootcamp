package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

// se crea un un struct mockConfig
type mockConfig struct {
	clienteUsado string
}

// El mock debe implementar el método necesario y comprobar que SumaEnabled sea llamado y que se haga exactamente con el mismo cliente que recibió SumarRestricted
func (m *mockConfig) SumaEnabled(cliente string) bool {
	m.clienteUsado = cliente
	return true
}
func TestSumarRestricted(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "John Doe"
	resultadoEsperado := 8
	// Se genera el objeto dummy a usar para satisfacer la necesidad de la función Sumar
	myMock := &mockConfig{}
	// Se ejecuta el test y se valida el resultado y que el mock haya registrado la información correcta
	resultado := SumarRestricted(num1, num2, myMock, cliente)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	assert.Equal(t, cliente, myMock.clienteUsado)
}
