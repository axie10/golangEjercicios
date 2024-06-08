package main

import (
	"fmt"
	// "strconv"
	"math"
)

// //las constantes si queremos no se usan no nos va a dar error
// const(
// 	// Pi = 3.14
// 	Lunes = iota + 1
// 	Martes
// 	Miercoles
// 	Jueves
// 	Viernes
// 	Sabado
// 	Domingo
// )

func main() {

	// 	const edad = 10
	// 	var edad1 int = edad + 10

	// 	var (
	// 		name = "Golang"
	// 		age int8 = 10
	// 		x, y, z int8 = 1, 2, 3
	// 	)

	// 	fmt.Println("Hello World",age, name, x, y, z)
	// 	fmt.Println(Martes)
	// 	fmt.Println(edad1)
	// 	// fmt.Println(Pi)

	/*TIPOS DE DATOS*/

	// var edad int8 = 10
	// var decimal float32 = 3.14
	// var a byte = 'a'
	// s := "hola"
	// var r rune = '𬡜'

	// fmt.Println(edad)
	// fmt.Println(decimal)
	// fmt.Println(math.MinInt32)
	// fmt.Println(a)
	// fmt.Println(s[0])
	// fmt.Println(r)

	/*VAROLES POR DEFECTO DE LOS TIPOS DE DATOS*/
	// var (
	// 	defaulint int
	// 	defaultuint uint
	// 	defaultfloat float32
	// 	defaulbool bool
	// 	defaultstring string
	// )

	// fmt.Println(defaulbool, defaulint, defaultfloat, defaultuint, defaultstring)

	/*CONVERSIONES DE TIPOS DE DATOS*/
	// var num16 int16 = 50
	// var num32 int32 = 100
	// var numfloat float32 = 56.56
	// fmt.Println(int32(num16) + num32)
	// fmt.Println(num16 + int16(num32))
	// fmt.Println(int32(numfloat))

	// //realizar conversion de string a int i es el numero y b es el error
	// s:= "56"
	// i,b := strconv.Atoi(s)
	// fmt.Println(i + 2)
	// fmt.Println(b)

	// //vamos a convertir del tipo int a string
	// m := 45
	// p := strconv.Itoa(m)
	// fmt.Println(p)

/*PAQUETE FMT*/

	// fmt.Print("hola Print")
	// fmt.Println("hola mundo")

	// name := "pepe"
	// age := 34
	// fmt.Printf("hola me llamo %s y tengo %d años\n", name , age)
	// //para guardar en una variable una frase y lo formatea todo a string 'fmt.Sprintf()' 
	// greeting := fmt.Sprintf("hola me llamo %s y tengo %d años", name , age)
	// fmt.Println(greeting)

	// //vamos a hacer que el usuario ingrese las variables con 'fmt.Scanln()'
	// var name1 string
	// var age1 int
	// var apellido string
	// fmt.Print("Ingrese su nombre: ")
	// fmt.Scanln(&name1, &apellido)
	// fmt.Print("Ingrese su edad: ")
	// fmt.Scanln(&age1)
	// //usamos %s patra string y %d para int
	// fmt.Printf("Hola me llamo %s %s y tengo %d años\n", name1,apellido, age1)
	// //el %v lo usamos para cuando no sabemos que tipo de dato vamos a introducir
	// fmt.Printf("Hola me llamo %v %v y tengo %v años\n", name1,apellido, age1)
	// //si ponemos %T nos dira el tipo de dato que queremos mostrar
	// fmt.Printf("El tipo de name es %T:\n", name1)
	// fmt.Printf("El tipo de age es %T:\n", age)

/*OPERACIONES ARITMETICAS*/

	// a := 23
	// b := 3
	// b--
	// b = b -7
	// a += 5
	// c := math.Pow(2 , 3)
	// fmt.Println(c)
	// fmt.Println(b)
	// fmt.Println(a)
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)
	// fmt.Println(math.Pi)
	// fmt.Println(math.Sqrt(64))
	// fmt.Println(math.Cbrt(27))

	//variables
	var altura float64
	var base float64
	const precision = 2

	//datos ingresados por el usuario
	fmt.Println("Introduce la altura del triangulo")
	fmt.Scanln(&altura)
	fmt.Println("Introduce la base del triangulo")
	fmt.Scanln(&base)

	//hacemos los calculos
	hipotenusa := math.Sqrt(math.Pow(base,2) + math.Pow(altura,2))
	areaTriangulo := (base * altura) / 2
	perimetro := base + altura + hipotenusa

	//salida de datos, tambien aprendemos a formatear datos
	fmt.Printf( "La hipotenusa del triangulo es: %.*f\n", precision, hipotenusa)
	fmt.Printf( "El area del triangulo es: %.*f\n", precision ,areaTriangulo)
	fmt.Printf("El perimetro del triangulo es: %.*f\n",precision,perimetro)



}
