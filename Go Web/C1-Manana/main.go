package main

/*
Para utilizar Gin se requiere la versión 1.13+ de Go, una vez instalada, utilizamos el siguiente
comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Ejercicio 3")

	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/message
	router.GET("/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Peter 8-)",
		})
	})

	type product struct {
		Name      string
		Price     int
		Published bool
	}

	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`

	var p product

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	router.GET("/productos", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": p,
		})
	})

	router.GET("/GetAll", func(c *gin.Context) {

		datosBytes, err := ioutil.ReadFile("./productos.json")
		if err != nil {
			log.Fatal(err)
		}
		datosString := string(datosBytes)

		fmt.Println(datosString)

		c.JSON(200, gin.H{
			"message": datosString,
		})
	})

	router.Run()

}

//Ejecutamos en Postman

//localhost:8080/message

//localhost:8080/productos
