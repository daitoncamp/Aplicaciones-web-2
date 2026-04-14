package main

import "fmt"

func main() {

	//Ejercicio 1 "Mi Ficha Personal"
	//var name string = "Daiton Campuzano"
	//var edad int = 22
	//var carrera string = "Tecnologías de la información"
	//var semestre int = 6
	//var promedio float64 = 8.9

	//fmt.Printf("Hola, soy %s tengo %d años\n  Estudio la carrera de %s\n  Actualmente me encuentro en %d nivel y mi promedio es de %2f ", name, edad, carrera, semestre, promedio)

	//Ejercicio 2 "Operaciones con Datos"
	//const producto1 = 40
	//const producto2 = 30
	//const producto3 = 50

	//const total = producto1 + producto2 + producto3
	//const promedio = total / 3
	//const descuento = total * 0.15

	//fmt.Println("Precio del producto 1: $ ", producto1, "\nPrecio del producto 2: $ ", producto2, "\nPrecio del producto 3: $ ", producto3, "\nTotal: $", total, "\nPromedio: $", promedio, "\nTotal con 15% descuento: $", descuento)

	//Ejercicio 3 "Mini-Interacción (Bonus)"
	var nombre string
	var nota1, nota2 float64
	var promedio float64

	fmt.Print("¿Cómo te llamas? ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingresa tu nota 1: ")
	fmt.Scanln(&nota1)

	fmt.Print("Ingresa tu nota 2: ")
	fmt.Scanln(&nota2)

	promedio = (nota1 + nota2) / 2

	fmt.Printf("%s, tu promedio es: %.2f\n", nombre, promedio)

	if promedio >= 7 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}

	//  fmt.Println("¡Hola, Aplicaciones Web II!")

}
