package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// Es un slice de parámetros dados por la URL.
//Estos “Param” ocupan el mismo orden que en la URL.
type Param struct {
	Key   string
	Value string
}

//Definición de “Params”
type Params []Param

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

var empleados = map[string]string{
	"644": "Empleado A",
	"755": "Empleado B",
	"777": "Empleado C",
}

type Producto struct {
	Id              int     `form:"Id" json:"id"`
	Nombre          string  `form:"Nombre" json:"nombre"`
	Precio          float64 `form:"Precio" json:"precio"`
	Stock           int     `form:"Stock" json:"stock"`
	Codigo          string  `form:"Codigo" json:"codigo"`
	Publicado       bool    `form:"Publicado" json:"publicado"`
	FechaDeCreacion string  `form:"FechaDeCreacion" json:"fechaDeCreacion"`
}

func main() {
	server := gin.Default()
	server.GET("/", PaginaPrincipal)
	server.GET("/empleados/:id", BuscarEmpleado)
	server.GET("/empleadosparams", DevolverEmpleadoParams)

	server.GET("/productosparams", DevolverProductoParams)

	server.GET("/filtrarproductos", FiltrarProductosFecha)

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
		//ctxt.String(200, "Información del empleado %s, nombre: %s", ctxt.Param("id"), empleado)
		emple := Empleado{
			Nombre: empleado,
			Id:     ctxt.Param("id"),
		}

		ctxt.JSON(200, gin.H{
			"empleado": emple,
		})
	} else {
		ctxt.String(404, "Información del empleado ¡No existe!")
	}
}

func DevolverEmpleadoParams(ctxt *gin.Context) {

	var empleado Empleado
	//Nuestro objetivo aquí es asignar los campos de nuestra estructura con los datos que recibimos del request.
	if ctxt.BindQuery(&empleado) == nil {
		log.Println("====== Bind Por Query String ======")
		log.Println(empleado.Id)
		log.Println(empleado.Nombre)
		//ctxt.String(200, "Empleado: %s, Id: %s\n", empleado.Nombre, empleado.Id)
		ctxt.JSON(200, gin.H{
			"Empleado": empleado,
		})
	}
}

func DevolverProductoParams(ctxt *gin.Context) {

	var producto Producto
	//Nuestro objetivo aquí es asignar los campos de nuestra estructura con los datos que recibimos del request.
	if ctxt.BindQuery(&producto) == nil {
		log.Println(producto)

		ctxt.JSON(200, gin.H{
			"Producto": producto,
		})
	}
}

//Esta función solo mostrará aquellos empleados activos o inactivos, dependiente del parámetro active.
/*func FiltrarEmpleados(ctxt *gin.Context) {
	empleados := GenerarListaEmpleados()
	var filtrados []*Empleado
	for i, e := range empleados {
		if ctxt.Query("active") == e.Activo {
			filtrados = append(filtrados, &e)
		}
	}
   }*/

//Esta función solo mostrará aquellos productos filtrados por fechas
func FiltrarProductosFecha(ctxt *gin.Context) {

	datosBytes, err := ioutil.ReadFile("./productos.json")
	if err != nil {
		log.Fatal(err)
	}
	datosString := string(datosBytes)

	fmt.Println(datosString)
	var listProducts []Producto

	if err := json.Unmarshal([]byte(datosString), &listProducts); err != nil {
		log.Fatal(err)
	}

	fmt.Println(listProducts)

	productos := listProducts
	var filtrados []Producto
	for _, e := range productos {
		if ctxt.Query("fechadecreacion") == e.FechaDeCreacion {
			filtrados = append(filtrados, e)
		}
	}

	ctxt.JSON(200, filtrados) //Para devolver directamente todo el arreglo de Productos filtrados
}
