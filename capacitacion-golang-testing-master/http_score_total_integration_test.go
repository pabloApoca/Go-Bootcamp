package http_score_total

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{Store: store}
	engine := server.GetMainEngine()
	player := "Beth"

	engine.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	engine.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	engine.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, "2", response.Body.String()) //Se hicieron 2 Post
}
