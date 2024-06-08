package main

import (
	// "errors"
	"fmt"
	// "os"
	// "fmt"
	// "strconv"
)

//funcion con manejo de errores personalizado
// func divide(diviendo, divisor int)(int, error){
// 	if divisor == 0 {
// 		return 0, fmt.Errorf("No se puede dividir entre 0")
// 	}
// 	return diviendo / divisor, nil
// }

//funcion con PANIC
func divide(diviendo, divisor int){
	//funxion anonima
	defer func (){
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	// validateZero(divisor)
	fmt.Println(diviendo/divisor)
}

// func validateZero(divisor int){
// 	if divisor == 0{
// 		panic("No se puede dividir por cero")
// 	}
// }






func main(){

/*MANEJO DE ERRORES POR DEFECTO QUE DEVUELVE GO Y MANEJO DE ERRORES PERSONALIZADOS QUE HACEMOS NOSTRSO EN UNA FUNCION PROPIA*/
	// str := "1234f"

	// num, err := strconv.Atoi(str)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// } 
	// fmt.Println("Numero:", num)

	// result, error := divide(2,0)

	// if error != nil {
	// 	fmt.Println("Error:",error)
	// 	return
	// }

	// fmt.Println("Resultado:",result)

/*DEFER 
cuando se pone en una funcion lo que se aplica es que se ejecuta al final de todo*/
	// defer fmt.Println(3)
	// fmt.Println(1)
	// fmt.Println(2)

	// file, err := os.Create("hola.txt")
	// defer file.Close()

	// if err != nil {
	// 	fmt.Println("Error0:", err)
	// 	return
	// }

	// _, err1 := file.Write([]byte("Hola programadores"))
	// if err1 != nil {
	// 	fmt.Println("Error1:", err1)
	// 	return
	// }

/*PANIC & RECOVER se reserva para casos excepcionales para el manejo de errores normales es con el err que devuelven las funciones file, err := func(){}
se utiliza para manejar situaciones excepcionales o manejar errores graves
RECOVER sirve para controlar el panic
No hace falta hacer la panic manual, como con la funcion validatezero para usar el recover, pero nos sacara otro texto en el error, sin embargo si manejamos el panic nos sacara el texto del error que nosotros queremos*/

	//al utilizar el defer hacemos que se ejcute todo y por ultimo nos saque el error y con el recover capturamos el error y lo manejamos de mejor manera captura el panic y solo no saca la frase del error, sin embargo si dejamos solo el panic nos sacara tambien en las lineas en las que se esta produciendo el error
	divide(100,10)
	divide(100,20)
	divide(100,0)
	divide(100,300)
	


}