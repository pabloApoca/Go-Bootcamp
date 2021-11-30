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

type product struct {
	Name      string
	Price     int
	Published bool
}

type Productos struct {
	Id              int     `form:"Id" json:"id"`
	Nombre          string  `form:"Nombre" json:"nombre"`
	Precio          float64 `form:"Precio" json:"precio"`
	Stock           int     `form:"Stock" json:"stock"`
	Codigo          string  `form:"Codigo" json:"codigo"`
	Publicado       bool    `form:"Publicado" json:"publicado"`
	FechaDeCreacion string  `form:"FechaDeCreacion" json:"fechaDeCreacion"`
}

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

	//Devuelvo un producto
	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`
	var p product
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	router.GET("/producto", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": p,
		})
	})

	//Devuelvo la lista de productos
	router.GET("/GetAll", func(c *gin.Context) {

		datosBytes, err := ioutil.ReadFile("./productos.json")
		if err != nil {
			log.Fatal(err)
		}
		datosString := string(datosBytes)

		fmt.Println(datosString)
		var listProducts []Productos

		if err := json.Unmarshal([]byte(datosString), &listProducts); err != nil {
			log.Fatal(err)
		}

		fmt.Println(p)

		/*c.JSON(200, gin.H{ //Para devolverlo de la forma clave valor
			"listProducts": listProducts,
		})*/

		c.JSON(200, listProducts) //Para devolver directamente todo el arreglo de Productos

	})

	router.Run()

}

//Ejecutamos en Postman

//localhost:8080/message

//localhost:8080/productos
