package main

/*
Comenzamos con:
go mod init C2-Tarde
Para agregar  a nuestra carpeta mod

Para utilizar Gin se requiere la versión 1.13+ de Go, una vez instalada, utilizamos el siguiente
comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/
import (
	"github.com/gin-gonic/gin"
	"C2-Tarde\cmd\server\handler"
	"C2-Tarde\internal\products"
 )

//Instanciamos cada capa del dominio Productos y utilizaremos los métodos del controlador para cada endpoint
func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)
 
	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
 
}
