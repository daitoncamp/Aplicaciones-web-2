package cafeteria

import (
	"errors"
)

// ERRORES EXPORTADOS
// -----------------------------------------------------------------------------

var (
	ErrClienteNoEncontrado  = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente    = errors.New("stock insuficiente")
	ErrSaldoInsuficiente    = errors.New("saldo insuficiente del cliente")
)

// Cliente
type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

// Categoria
type Categoria struct {
	ID     int
	Nombre string
}

// Producto

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria Categoria
}

// Pedido
type Pedido struct {
	ID       int
	Cliente  Cliente
	Producto Producto
	Cantidad int
	Total    float64
	Fecha    string
}

// INTERFAZ REPOSITORIO
type Repository interface {
	//clientes
	GuardarCliente(cliente Cliente) error
	ObtenerCliente(id int) (Cliente, error)
	ListarClientes() []Cliente

	//productos
	GuardarProducto(producto Producto) error
	ObtenerProducto(id int) (Producto, error)
	ListarProductos() []Producto
}

// RepoMemoria
type RepoMemoria struct {
	clientes  []Cliente
	productos []Producto
	pedidos   []Pedido
}

// Guardar clientes
func (r *RepoMemoria) GuardarCliente(c Cliente) error {
	r.clientes = append(r.clientes, c)
	return nil
}

// Obetener clientes
func (r *RepoMemoria) ObtenerCliente(id int) (Cliente, error) {
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}

// Listar clientes
func (r *RepoMemoria) ListarClientes() []Cliente {
	return r.clientes
}

// Guardar producto
func (r *RepoMemoria) GuardarProducto(p Producto) error {
	r.productos = append(r.productos, p)
	return nil
}

// Obtener producto
func (r *RepoMemoria) ObtenerProducto(id int) (Producto, error) {
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}

// Listar prodctos
func (r *RepoMemoria) ListarProductos() []Producto {
	return r.productos
}

// Constructor
func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{}
}

// Verificación en tiempo de compilación:
// Si RepoMemoria NO cumple Repository, esto da error al compilar.

var _ Repository = (*RepoMemoria)(nil)
