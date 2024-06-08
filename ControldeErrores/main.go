package main

import (
	// "errors"
	"fmt"
	"os"
	// "fmt"
	// "strconv"
)

//funcion con manejo de errores personalizado
func divide(diviendo, divisor int)(int, error){
	if divisor == 0 {
		return 0, fmt.Errorf("No se puede dividir entre 0")
	}
	return diviendo / divisor, nil
}






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

/*PANIC & RECOVER
se utiliza para manejar situaciones excepcionales o manejar errores graves*/

	





}