package main

import (
	"errors"
	"fmt"
	"semana3_vivo/internal/cafeteria"
)

func main() {

	/*c := cafeteria.Cliente{
				ID:     1,
				Nombre: "Ana",
			}

			fmt.Println(c)

	}*/

	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()

	// Guardar clientes3 con su categoría anidada
	repo.GuardarCliente(cafeteria.Cliente{})

	// Clientes
	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Juan", Carrera: "TI", Saldo: 20})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Ana", Carrera: "Civil", Saldo: 15})

	// Categorías
	bebidas := cafeteria.Categoria{ID: 1, Nombre: "Bebidas"}
	snacks := cafeteria.Categoria{ID: 2, Nombre: "Snacks"}

	// Productos (con categoría ANIDADA)
	repo.GuardarProducto(cafeteria.Producto{
		ID: 1, Nombre: "Cola", Precio: 1.0, Stock: 10, Categoria: bebidas,
	})
	repo.GuardarProducto(cafeteria.Producto{
		ID: 2, Nombre: "Jugo", Precio: 1.2, Stock: 8, Categoria: bebidas,
	})
	repo.GuardarProducto(cafeteria.Producto{
		ID: 3, Nombre: "Pan", Precio: 0.5, Stock: 20, Categoria: snacks,
	})

	//Obtener cliente existente
	fmt.Println("\n--- Buscar cliente existente ---")
	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}

	//Cliente que no existe
	fmt.Println("\n--- Buscar cliente inexistente ---")
	_, err = repo.ObtenerCliente(99)
	if err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, cafeteria.ErrClienteNoEncontrado) {
			fmt.Println("→ Error confirmado: cliente no encontrado")
		}
	}

	//Listar productos
	fmt.Println("\n--- Lista de productos ---")
	for _, p := range repo.ListarProductos() {
		fmt.Printf("[%d] %s - $%.2f - Stock: %d - Categoria: %s\n",
			p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria.Nombre)
	}

	//Mostrar Pedido
	fmt.Println("\n--- Pedidos ---")

	cliente, _ := repo.ObtenerCliente(1)
	producto, _ := repo.ObtenerProducto(1)

	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  cliente,
		Producto: producto,
		Cantidad: 2,
		Total:    producto.Precio * 2,
		Fecha:    "2026-04-23",
	}

	fmt.Println("Pedido ID:", pedido.ID)
	fmt.Println("Cliente:", pedido.Cliente.Nombre)
	fmt.Println("Producto:", pedido.Producto.Nombre)
	fmt.Println("Categoria:", pedido.Producto.Categoria.Nombre)
	fmt.Println("Total:", pedido.Total)
}
