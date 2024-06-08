package main

import (
	"fmt"
	// "internal/syscall/windows"
	// "runtime"
	// "time"
)

func main() {

	/*IF ELSE*/
	// var numero1 int
	// var numero2 int
	// numero1 = 67

	// if(numero1 == 67){
	// 	numero2 = 1
	// 	fmt.Println(numero2)
	// } else{
	// 	numero2 = 2
	// 	fmt.Println(numero2)
	// }

	//SE PUEDE HACER ASI
	// t := time.Now()
	// hora := t.Hour()

	// if hora < 12 {
	// 	fmt.Println("Mañana")
	// } else if hora < 19{
	// 	fmt.Println("tarde")
	// } else {
	// 	fmt.Println("noche")
	// }

	//TAMBIEN SE PUEDE DECLARAR LA VARIABLE DENTRO DE LA CONDICION
	// if t := time.Now(); t.Hour() < 12 {
	// 	fmt.Println("Mañana")
	// } else if t.Hour() < 19{
	// 	fmt.Println("tarde")
	// } else {
	// 	fmt.Println("noche")
	// }

	/*SWITCH*/

	// os := runtime.GOOS

	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("windows")

	// case "linux":
	// 	fmt.Println("linux")

	// case "darwin":
	// 	fmt.Println("mac")
	// default:
	// 	fmt.Println("Otro SO")
	// }

	// switch t := time.Now(); {
	// case t.Hour() < 12:
	// 	fmt.Println("Mañana")
	// case t.Hour() > 12:
	// 	fmt.Println("Tarde")
	// case t.Hour() > 19:
	// 	fmt.Println("Noche")
	// default:
	// 	fmt.Println("Es otra hora")
	// }

	/*BUCLE FOR*/
	// var i int

	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 1; i < 10; i++ {
	// 	// fmt.Println(i)
	// 	if i == 7 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	num(7, 5)
	hello("Jose Luis")
	hello2("Jose andres")
	saludo := hello2("Jose andres")
	fmt.Println(saludo)
	suma, resta := cal(5, 3)
	fmt.Println("esta es la suma", suma)
	fmt.Println("esta es la resta", resta)

	sum, rest := cal2(5, 3)
	fmt.Println("esta es la suma", sum)
	fmt.Println("esta es la resta", rest)

}

func num(x int, y int) {
	if x > y {
		fmt.Println("x es mayor que y")
	} else {
		fmt.Println("y es mayor que x")
	}
}

func hello(name string) {
	fmt.Println("Hello", name)
}

func hello2(name string) string {
	return "Hello" + name
}

func cal(a, b int) (int, int) {
	suma := a + b
	resta := a - b
	return suma, resta
}

func cal2(a, b int) (sum, rest int) {
	sum = a + b
	rest = a - b
	return
}
