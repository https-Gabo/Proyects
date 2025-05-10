package main

import (
	"fmt"
)

func AgregarProducto(inventario map[string]int, producto string, cantidad int) {
	inventario[producto] += cantidad
	fmt.Printf("Producto '%s' añadido/actualizado. Cantidad: %d\n", producto, inventario[producto])
}

func VerCantidadProducto(inventario map[string]int, producto string) int {
	return inventario[producto]
}

func ListarInventario(inventario map[string]int) {
	if len(inventario) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}
	fmt.Println("Inventario:")
	for producto, cantidad := range inventario {
		fmt.Printf("- %s: %d\n", producto, cantidad)
	}
}

func EliminarProducto(inventario map[string]int, producto string) {
	delete(inventario, producto)
	fmt.Printf("Producto '%s' eliminado del inventario.\n", producto)
}

func main() {
	inventario := make(map[string]int)

	var opcion int
	for {
		fmt.Println("\nMenú de gestión de inventario:")
		fmt.Println("1. Agregar producto")
		fmt.Println("2. Ver cantidad de un producto")
		fmt.Println("3. Listar inventario")
		fmt.Println("4. Eliminar producto")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var producto string
			var cantidad int
			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scanln(&producto)
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scanln(&cantidad)
			AgregarProducto(inventario, producto, cantidad)

		case 2:
			var producto string
			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scanln(&producto)
			cantidad := VerCantidadProducto(inventario, producto)
			fmt.Printf("Cantidad de '%s': %d\n", producto, cantidad)

		case 3:
			ListarInventario(inventario)

		case 4:
			var producto string
			fmt.Print("Ingrese el nombre del producto a eliminar: ")
			fmt.Scanln(&producto)
			EliminarProducto(inventario, producto)

		case 5:
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}
