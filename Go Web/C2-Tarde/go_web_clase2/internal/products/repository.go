package products

<<<<<<< HEAD:Go Web/C2-Tarde/go_web_clase2/internal/products/repository.go
//Empezaremos con implementar el repositorio. Lo primero que haremos es declarar las entidades que
//utilizaremos: Producto, Productos y el ultimo ID almacenado.
=======
//Empezaremos con implementar el repositorio. Lo primero que haremos es declarar las entidades que utilizaremos:
//Producto, Productos y el ultimo ID almacenado.

//import "C2-Tarde/internal/products"
import "C2-Tarde/internal/products"

>>>>>>> 52c60fd8ce5b5161a4b0445283479977635c1173:Go Web/C2-Tarde/internal/products/repository.go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product
var lastID int

<<<<<<< HEAD:Go Web/C2-Tarde/go_web_clase2/internal/products/repository.go
//Implementaremos la interface Repositorio con sus métodos y una función que nos devuelva el
//repositorio que se utilizará.
=======
//Implementaremos la interface Repositorio con sus métodos y una función que nos devuelva el repositorio que se utilizará.
>>>>>>> 52c60fd8ce5b5161a4b0445283479977635c1173:Go Web/C2-Tarde/internal/products/repository.go
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
}

type repository struct{} //f2

func NewRepository() Repository {
	return &repository{}
}

//Implementaremos los métodos:
//GetAll: Obtener todos los Productos.
//LastID: Obtener el ultimo ID almacenado.
<<<<<<< HEAD:Go Web/C2-Tarde/go_web_clase2/internal/products/repository.go

=======
>>>>>>> 52c60fd8ce5b5161a4b0445283479977635c1173:Go Web/C2-Tarde/internal/products/repository.go
func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

//Implementaremos los métodos:
//Store: Guardará la información de producto, asignará el último ID a la variable y nos retorna la entidad Producto.
func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	p := Product{id, nombre, tipo, cantidad, precio}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
