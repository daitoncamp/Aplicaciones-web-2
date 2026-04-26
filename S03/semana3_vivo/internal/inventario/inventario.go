package inventario

import (
	"errors"
	"fmt"
)

var (
	ErrCategoriaNoEncontrada = errors.New("categoria no encontrada")
	ErrProductoNoEncontrada  = errors.New("producto no encontrada")
)

type Repositorio interface {
	Guardar(p Producto) error
	BuscarProductoPorId(id int) (Producto, error)
	ListarProdcutos() []Producto
}

type RepoMemoria struct {
	productos []Producto
}

func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{productos: []Producto{}}
}

type Categoria struct {
	ID     int
	Nombre string
}

type Producto struct {
	ID          int
	Nombre      string
	Precio      float64
	Stock       int
	CategoriaID int // referencia a Categoria por ID, NO anidación
}

// -----------------------------------------------------------------------------
// "BASE DE DATOS" EN MEMORIA
// -----------------------------------------------------------------------------

var categorias = []Categoria{}
var productos = []Producto{}

// -----------------------------------------------------------------------------
// FUNCIONES DE CATEGORÍAS
// -----------------------------------------------------------------------------

func AgregarCategoria(c Categoria) {
	categorias = append(categorias, c)
}

func BuscarCategoriaPorID(id int) (Categoria, error) {
	// PROBLEMA PEDAGÓGICO: si no existe, devolvemos Categoria{} (zero value).
	// No hay forma de que el llamador sepa si existía o no. En Semana 3
	// esto se arregla retornando (Categoria, error).
	for _, c := range categorias {
		if c.ID == id {
			return c, nil
		}
	}
	return Categoria{}, ErrCategoriaNoEncontrada
}

// -----------------------------------------------------------------------------
// FUNCIONES DE PRODUCTOS
// -----------------------------------------------------------------------------

func (r *RepoMemoria) Guardar(p Producto) error {
	r.productos = append(r.productos, p)
	return nil
}

func (r *RepoMemoria) BuscarProductoPorID(id int) (Producto, error) {
	// Mismo problema que BuscarCategoriaPorID.
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrada
}

func ListarProductos() {
	fmt.Println("=== LISTADO DE PRODUCTOS ===")
	for _, p := range productos {
		cat, err := BuscarCategoriaPorID(p.CategoriaID)
		if err != nil {
			fmt.Printf("error al buscar la catgoria %s\n", p.Nombre)
			continue
		}
		fmt.Printf("[%d] %s — $%.2f — stock: %d — categoría: %s\n",
			p.ID, p.Nombre, p.Precio, p.Stock, cat.Nombre)
	}
}

func CalcularValorInventario() float64 {
	total := 0.0
	for _, p := range productos {
		total += p.Precio * float64(p.Stock)
	}
	return total
}
