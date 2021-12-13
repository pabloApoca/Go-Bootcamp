package main

import (
	http_score_total "github.com/IgnacioFalco/capacitacion-golang-testing"
)

func main() {

	server := &http_score_total.PlayerServer{Store: http_score_total.NewInMemoryPlayerStore()}
	server.ServeHTTP()
}
