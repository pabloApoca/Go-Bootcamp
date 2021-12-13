package http_score_total

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	ObtainPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) GetMainEngine() *gin.Engine {

	r := gin.Default()
	players := r.Group("/players")
	players.POST("/:name", func(c *gin.Context) {
		p.addScore(c, c.Param("name"))

	})
	players.GET("/:name", func(c *gin.Context) {
		p.showScore(c, c.Param("name"))

	})
	fmt.Println("Server Running in http://localhost:8080")
	return r
}

func (p *PlayerServer) ServeHTTP() {
	p.GetMainEngine().Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (p *PlayerServer) addScore(c *gin.Context, name string) {
	p.Store.RecordWin(name)
	c.String(http.StatusOK, "OK") //este http status deberia ser otro (StatusOK)
}
func (p *PlayerServer) showScore(c *gin.Context, name string) {

	score := p.Store.ObtainPlayerScore(name)
	if score == 0 {
		c.String(http.StatusNotFound, "NOT_FOUND")
		return
	}
	c.String(http.StatusOK, strconv.Itoa(score))
}
