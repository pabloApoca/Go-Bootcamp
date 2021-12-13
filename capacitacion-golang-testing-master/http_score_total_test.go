package http_score_total

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) ObtainPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestHttpScore(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Rick":  20,
			"Morty": 35,
		}, nil,
	}
	server := &PlayerServer{Store: &store}
	engine := server.GetMainEngine()

	t.Run("Devuelve el puntaje de Rick", func(t *testing.T) {
		req := newGetScoreRequest("Rick")
		res := httptest.NewRecorder()
		engine.ServeHTTP(res, req)
		assertStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, "20", res.Body.String()) // El entero debe ser 20

	})

	t.Run("Retorna 404 cuando no encuentra al jugador", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest("Pablo") //El nombre debe ser distinto o incorrecto
		engine.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{}, nil,
	}
	server := &PlayerServer{Store: &store}
	engine := server.GetMainEngine()

	t.Run("Registra el puntaje del jugador cuando se hace un POST", func(t *testing.T) {
		player := "Jerry"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		engine.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK) //El codigo principal deberia ser statusOK

		if len(store.winCalls) != 1 {
			t.Errorf("Se obtuvieron %d llamadas a RecordWin. Se esperaba: %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("No se registró correctamente. El ganador obtuvo '%s', y se esperaban '%s'", store.winCalls[0], player)
		}
	})
}

func assertStatus(t *testing.T, result int, expected int) {
	if result != expected {
		t.Errorf("Se obtuvo otro estado de error: '%d'. Se esperaba '%d'", result, expected)
	}
}

func assertResponseBody(t *testing.T, expected string, result string) {
	if result != expected {
		t.Errorf("Resultado: '%s'. Esperado '%s'", result, expected)
	}
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}
