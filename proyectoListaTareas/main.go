package main

import (
	"bufio"
	"fmt"
	"os"
)

/*Interface de tareas*/
type Tarea struct {
	titulo      string
	descripcion string
	estado      bool
}

// interface tareas
type Tareas struct {
	tarea []Tarea
}

// Metodo para agregar tarea
func (l *Tareas) addTarea(t Tarea) {
	l.tarea = append(l.tarea, t)
}

// Metodo para marcar como completada la tarea
func (l *Tareas) completeTarea(index int) {
	l.tarea[index].estado = true
}

// Metodo para editar tarea
func (l *Tareas) editTarea(index int, t Tarea) {
	l.tarea[index] = t
}

// Metodo para eliminar tarea
func (l *Tareas) deleteTarea(index int) {
	l.tarea = append(l.tarea[:index], l.tarea[index+1:]...)
}

func main() {

	listadoTareas := Tareas{}

	//instancia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Selecciona una opcion:")
		fmt.Println("Añadir tarea => 1\nEditar tarea => 2\nEliminar tarea => 3\nCompletar tarea => 4\nSalir => 5")
		fmt.Print("Ingrese una opcion:")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//variable
			var t Tarea
			//entrada de datos
			fmt.Println("Añaded titulo:")
			t.titulo, _ = leer.ReadString('\n')
			fmt.Println("Añaded descripcion:")
			t.descripcion, _ = leer.ReadString('\n')
			//metodo añadir
			listadoTareas.addTarea(t)
		case 2:
			//variables
			var t Tarea
			var index int
			//entrada de datos
			fmt.Println("Añaded titulo:")
			t.titulo, _ = leer.ReadString('\n')
			fmt.Println("Añaded descripcion:")
			t.descripcion, _ = leer.ReadString('\n')
			fmt.Print("Que tarea quieres editar:")
			fmt.Scanln(&index)
			//metodo de editar
			listadoTareas.editTarea(index, t)
		case 3:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desee eliminar:")
			fmt.Scanln(&index)
			listadoTareas.deleteTarea(index)
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desee completar:")
			fmt.Scanln(&index)
			listadoTareas.completeTarea(index)
		case 5:
			fmt.Print("Saliendo del programa...")
			return
		default:
			fmt.Print("Opcion invalida")
		}

		fmt.Println("Listado de Tareas:")
		fmt.Println("===========================")
		for i, t := range listadoTareas.tarea {
			fmt.Printf("Tarea:%v\ntitulo:%vdescripcion:%vestado:%v\n", i, t.titulo, t.descripcion, t.estado)
		}
		fmt.Println("===========================")

	}

	// var titulo string
	// var descripcion string
	// var estado bool = false
	// var tarea Tarea

	// var tareas []Tarea

	// fmt.Println("Introduce tarea:")
	// fmt.Println("Titulo:")
	// fmt.Scanln(&titulo)
	// fmt.Println("Descripcion:")
	// fmt.Scanln(&descripcion)
	// tarea.titulo = titulo
	// tarea.descripcion = descripcion
	// tarea.estado = estado

	// tareas = append(tareas, tarea)

	// fmt.Println(tareas)

}
