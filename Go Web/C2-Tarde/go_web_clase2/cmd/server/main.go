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
<<<<<<< HEAD:Go Web/C2-Tarde/go_web_clase2/cmd/server/main.go
//	"github.com/pabloApoca/go_web_clase2/cmd/server\handler"

=======

//El main se encontrara en el directorio cmd/server.
//Importamos las dependencias necesarias.
>>>>>>> 52c60fd8ce5b5161a4b0445283479977635c1173:Go Web/C2-Tarde/cmd/server/main.go
import (
	//"C2-Tarde/cmd/server/handler/products.go"
	"C2-Tarde/cmd/server/handler"
	"C2-Tarde/internal/products"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD:Go Web/C2-Tarde/go_web_clase2/cmd/server/main.go
	"github.com/pabloApoca/go_web_clase2/cmd/server/handler"
	"github.com/pabloApoca/go_web_clase2/internal/products"
=======
>>>>>>> 52c60fd8ce5b5161a4b0445283479977635c1173:Go Web/C2-Tarde/cmd/server/main.go
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

//Lo corremos con go run cmd/server/main.go
