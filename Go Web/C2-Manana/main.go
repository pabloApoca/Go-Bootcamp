package main

/*
Comenzamos con:
go mod init C2-Manana
Para agregar  a nuestra carpeta mod

Para utilizar Gin se requiere la versión 1.13+ de Go, una vez instalada, utilizamos el siguiente
comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

import "github.com/gin-gonic/gin"

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

var products []request
var lastID int

func main() {

	//Prueba de un Post basico
	/*r := gin.Default()
	r.POST("/productos", func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil { //ShouldBind: Si hay un error, gin no hace nada y deja en el desarrollador la responsabilidad de manipularlo.
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.ID = 4
		c.JSON(200, req)
	})
	r.Run()*/

	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/add", Guardar())
	r.Run()

}

func Guardar() gin.HandlerFunc {

	return func(c *gin.Context) { //Valido que el Token este bien
		token := c.GetHeader("token")

		if token != "eltokenqueapruebalaspeticiones" {
			c.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}

		var req request //Guado el producto
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Nombre == "" { //Valido el campo Nombre
			c.JSON(404, gin.H{
				"error": "el campo Nombre es requerido",
			})
			return
		}
		if req.Tipo == "" { //Valido el campo Tipo
			c.JSON(404, gin.H{
				"error": "el campo Tipo es requerido",
			})
			return
		}
		if req.Cantidad <= 0 { //Valido el campo Cantidad
			c.JSON(404, gin.H{
				"error": "el campo Cantidad es requerido",
			})
			return
		}
		if req.Precio <= 0 { //Valido el campo Precio
			c.JSON(404, gin.H{
				"error": "el campo Nombre es requerido",
			})
			return
		}

		lastID++
		req.ID = lastID

		products = append(products, req)

		c.JSON(200, req) //Devuelvo el producto
	}

}
