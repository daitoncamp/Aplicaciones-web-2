// =============================================================================
// SEMANA 3 - DÍA A - PUNTO DE PARTIDA DEL LIVE CODING
// =============================================================================
// Este es el estado del código al final de Semana 2 (variante con IDs).
// Todo vive en un único archivo: structs, slices, funciones y main.
//
// Durante la clase lo refactorizaremos a:
//   1. Paquetes separados (internal/inventario/)
//   2. Manejo de errores idiomático con (valor, error)
//   3. Interfaces como contrato entre paquetes
//
// Para ejecutar:   go run .
// =============================================================================

package main

import (
	"fmt"
	"semana3_vivo/inventario"
)

// -----------------------------------------------------------------------------
// STRUCTS (modelo con IDs — el tipo que luego verán en BD con GORM)
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------

func main() {
	// Cargamos datos iniciales
	inventario.AgregarCategoria(inventario.Categoria{ID: 1, Nombre: "Bebidas"})
	inventario.AgregarCategoria(inventario.Categoria{ID: 2, Nombre: "Snacks"})

	inventario.AgregarProducto(inventario.Producto{ID: 101, Nombre: "Agua 500ml", Precio: 0.50, Stock: 120, CategoriaID: 1})
	inventario.AgregarProducto(inventario.Producto{ID: 102, Nombre: "Cola 500ml", Precio: 1.00, Stock: 80, CategoriaID: 1})
	inventario.AgregarProducto(inventario.Producto{ID: 201, Nombre: "Papas fritas", Precio: 1.25, Stock: 45, CategoriaID: 2})

	// Usamos las funciones
	inventario.ListarProductos()

	fmt.Printf("\nValor total del inventario: $%.2f\n", inventario.CalcularValorInventario())

	// Buscamos un producto existente
	p, err := inventario.BuscarProductoPorID(101)
	if err != nil {
		fmt.Printf("error al buscar el producto %s\n", err.Error())

	}
	fmt.Printf("\nEncontrado: %s\n", p.Nombre)

	// Buscamos uno que NO existe — aquí está el problema que resolveremos
	fantasma, err := inventario.BuscarProductoPorID(999)
	if err != nil {
		fmt.Println("errpr al buscar producto \n", err)
	}
	fmt.Printf("Buscando ID 999: %+v\n", fantasma)
	// Output: Buscando ID 999: {ID:0 Nombre: Precio:0 Stock:0 CategoriaID:0}
	// ¿Cómo sabe el llamador si existía o no? No hay forma limpia.
}
