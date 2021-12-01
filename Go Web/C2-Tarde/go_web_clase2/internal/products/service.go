package products

<<<<<<< HEAD:Go Web/C2-Tarde/go_web_clase2/internal/products/service.go
//import repository
//import "github.com/pabloApoca/go_web_clase2/internal/products/repository.go"

//Implementaremos la interface Servicio con sus métodos y una función que reciba un Repositorio
//y nos devuelva el servicio que se utilizará, instanciado.
=======
//Implementaremos la interface Servicio con sus métodos y una función que reciba un Repositorio y nos devuelva
//el servicio que se utilizará, instanciado.
>>>>>>> 52c60fd8ce5b5161a4b0445283479977635c1173:Go Web/C2-Tarde/internal/products/service.go
type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

//Implementaremos el método GetAll que se encargará de pasarle la tarea al Repositorio y nos retorna un array de Productos.
func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

//El método Store se encargará de pasarle la tarea de obtener el ultimo ID y guardar el producto al Repositorio,
//el servicio se encargará de incrementar el ID.
func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, tipo, cantidad, precio)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}
