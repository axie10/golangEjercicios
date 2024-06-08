package main


import (
	"fmt"
)

/*CREAMOS LA ESTRUCTURA*/
type Persona struct{
	nombre string
	edad int16
	correo string
}

type Cartera struct {
	id string
	saldo int
	propietario string
}


func main () {

/*MATRICES, es una estructura de datos fija no se pueden agregar datos o quitar datos*/

//simples
	// var matriz = [...]int{10,20,30,40,50}

	// for i := 0; i < len(matriz); i++ {
	// 	fmt.Println(matriz[i])
	// }

	// for index, value := range matriz {
	// 	fmt.Println("Indice:", index, "Valor:", value)
	// }
	// for _, value := range matriz {
	// 	fmt.Println("Valor:", value)
	// }

//multidimensional

	// var matriz2 = [3][3] int {{1,2,3},{4,5,6},{7,8,9}}
	// fmt.Println(matriz2)

	// for i := 0; i < len(matriz2); i++ {
	// 	for v := 0; v < len (matriz2); v++ {
	// 		fmt.Println(matriz2[i][v])
	// 	}
	// }


/*REBANADA O SLICE, ademas de poder acceder a partes del elemento tambien permite agregar o quitar elementos*/

	// dias := []string {"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}

	// fmt.Println(dias)
	// diamenios := []string {"Sabado", "Domingo"}

	// dias = append(dias, diamenios...)
	// diasmenos := dias[0:2]

	// dias = append(dias[:4])



	// fmt.Println(dias)
	// fmt.Println(diasmenos)
	// fmt.Println(cap(dias))
	// fmt.Println(len(dias))

	//con make creamos un slice con una capacidad inicial
	// nombres := make([]int, 5)
	// nombres1 := []int{1,2,3,4,5}

	// copy(nombres, nombres1)
	// copy(nombres1, nombres)
	// fmt.Println(copy(nombres1, nombres))

	// fmt.Println(nombres)
	// fmt.Println(nombres1)

	// fmt.Println(nombres)
	// for i := 0; i < len(nombres); i++ {
	// 	nombres[i] = i 
	// }

	// fmt.Println(nombres)


	//que los multiplos de 3 salga fizz y que los multiplos de 2 salga buzz
	// for i := 1 ; i <= 100; i++ {

	// 	if i % 3 == 0 {
	// 		fmt.Println("fizz")
	// 	} else if i%5 == 0{
	// 		fmt.Println("buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	//es un numero primo?
	// for i := 1; i <= 100 ; i++ {

	// 	if (i % i == 1 && i % 1 == i) {
	// 		fmt.Println(i, "es un nuemro primo\n--------")
	// 	} else {
	// 		fmt.Println(i, "no es un numero primo")
	// 	}
	// }

/*MAPA*/

	// colors := map[string]string {
	// 	"rojo":"#FF0000",
	// 	"verde":"#00FF00",
	
	// fmt.Println(colors["azul"])
	//introducir valores a los mapas
	// colors["negro"] = "#000000"
	// fmt.Println(colors["negro"])
	//guardamos en una variable le valor de uno de las posiciones de los mapas y asdemas usamos el ok que es un booleano para ver si existe o no
	// valor,ok := colors["verde"]
	// if ok {
	// 	fmt.Println(valor)
	// } else {
	// 	fmt.Println("No existe la clave")
	// }

	// //borrar un elemento de un mapa
	// delete(colors, "rojo")
	// fmt.Println(colors)

	// //iteramos los valores del mapa
	// for clave, valor := range colors{
	// 	fmt.Printf("Clave: %v, Valor: %v\n", clave, valor)
	// }

/*ESTRUCTURAS*/

	// var p1 Persona
	// p1.nombre = "pepe"
	// p1.edad = 25
	// p1.correo = "pepe@gmail.com"

	// p2 := Persona {"roberto", 23, "roberto@gmail.com"}
	// p2.edad = 40

	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(p1.nombre)

/*PUNTEROS*/
//con punuteros podemos cambiar el valor de la memoria de una variable
	// var numero int = 10
	// // var p *int = &numero

	// // fmt.Println(numero)
	// // fmt.Println(&numero)
	// fmt.Println(numero)
	// pepe(&numero)
	// fmt.Println(numero)

	p  := Persona { "pepe",16,"pepe@gmail.com" }
	p.dihola()
	c := Cartera { "1", 3200, "jose"}
	c.obtenerDatos()



}

/*METODO*/
/*DE ESTA MANERA HACEMOS UN METODO, HACEMOS UNA FUNCION NORMAL Y LE PASAMOS ANTES DEL NOMBRE DE LA FUNCION UNA VARIABLE*/
func (p *Persona) dihola(){
	fmt.Println("Hola mi nombre es", p.nombre)
}

func (c *Cartera) obtenerDatos(){
	fmt.Printf("La cartera es de %v y tiene %v dinero", c.propietario, c.saldo)
}


func pepe (p *int){
	*p = 20
}