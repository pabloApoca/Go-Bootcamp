package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var empleados = map[string]string{
	"644": "Empleado A",
	"755": "Empleado B",
	"777": "Empleado C",
}

// Es un slice de parámetros dados por la URL.
//Estos “Param” ocupan el mismo orden que en la URL.

type Param struct {
	Key   string
	Value string
}

//Retorna el valor del “param” de la URL.
//Es un atajo para c.Params.ByName(key)
/*func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}*/

//Definimos nuestra estructura de información
type Empleado struct {
	// Una etiqueta de struct se cierra con caracteres de acento grave `

	Nombre string `form:"name" json:"name"`
	Id     string `form:"id" json:"id"`
	Activo string `form:"active" json:"activa" binding:"required"`
}

//Definición de “Params”
type Params []Param

func main() {
	server := gin.Default()
	server.GET("/", PaginaPrincipal)
	server.GET("/empleados/:id", BuscarEmpleado)
	server.GET("/empleadosparams", BuscarEmpleadoParams)

	server.Run(":8080")
}

//Este handler se encargará de responder a /.
func PaginaPrincipal(ctxt *gin.Context) {
	ctxt.String(200, "¡Bienvenido a la Empresa Gophers!")
}

//Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func BuscarEmpleado(ctxt *gin.Context) {
	empleado, ok := empleados[ctxt.Param("id")]
	if ok {
		ctxt.String(200, "Información del empleado %s, nombre: %s", ctxt.Param("id"), empleado)
	} else {
		ctxt.String(404, "Información del empleado ¡No existe!")
	}
}

func BuscarEmpleadoParams(ctxt *gin.Context) {

	var empleado Empleado
	//Nuestro objetivo aquí es asignar los campos de nuestra estructura con los datos que recibimos del request.
	if ctxt.BindQuery(&empleado) == nil {
		log.Println("====== Bind Por Query String ======")
		log.Println(empleado.Id)
		log.Println(empleado.Nombre)
		ctxt.String(200, "(Query String) - Empleado: %s, Id: %s\n", empleado.Nombre, empleado.Id)
	}
}
