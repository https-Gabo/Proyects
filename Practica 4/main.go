package main

import (
	"fmt"
)

type Estudiante struct {
	Nombre  string
	ID      int
	Carrera string
}

func AgregarEstudiante(slice []Estudiante, estudiante Estudiante) []Estudiante {
	return append(slice, estudiante)
}

func BuscarEstudiantePorID(slice []Estudiante, id int) *Estudiante {
	for i := range slice {
		if slice[i].ID == id {
			return &slice[i] 
		}
	}
	return nil 
}

func ListarEstudiantes(slice []Estudiante) {
	if len(slice) == 0 {
		fmt.Println("No hay estudiantes en el registro.")
		return
	}
	fmt.Println("Registro de estudiantes:")
	for _, estudiante := range slice {
		fmt.Printf("Nombre: %s, ID: %d, Carrera: %s\n", estudiante.Nombre, estudiante.ID, estudiante.Carrera)
	}
}

func main() {
	registroEstudiantes := []Estudiante{} 

	var opcion int
	for {
		fmt.Println("\nMenú de gestión de estudiantes:")
		fmt.Println("1. Agregar estudiante")
		fmt.Println("2. Buscar estudiante por ID")
		fmt.Println("3. Listar estudiantes")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1: 
			var nombre, carrera string
			var id int
			fmt.Print("Ingrese el nombre del estudiante: ")
			fmt.Scanln(&nombre)
			fmt.Print("Ingrese el ID del estudiante: ")
			fmt.Scanln(&id)
			fmt.Print("Ingrese la carrera del estudiante: ")
			fmt.Scanln(&carrera)

			nuevoEstudiante := Estudiante{Nombre: nombre, ID: id, Carrera: carrera}
			registroEstudiantes = AgregarEstudiante(registroEstudiantes, nuevoEstudiante)
			fmt.Println("Estudiante agregado correctamente.")

		case 2: 
			var id int
			fmt.Print("Ingrese el ID del estudiante a buscar: ")
			fmt.Scanln(&id)

			estudiante := BuscarEstudiantePorID(registroEstudiantes, id)
			if estudiante != nil {
				fmt.Printf("Estudiante encontrado: Nombre: %s, ID: %d, Carrera: %s\n", estudiante.Nombre, estudiante.ID, estudiante.Carrera)
			} else {
				fmt.Println("Estudiante no encontrado.")
			}

		case 3: 
			ListarEstudiantes(registroEstudiantes)

		case 4: 
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}
