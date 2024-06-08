package main

import (
	// "errors"
	"fmt"
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

	// str := "1234f"

	// num, err := strconv.Atoi(str)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// } 
	// fmt.Println("Numero:", num)

	result, error := divide(2,0)

	if error != nil {
		fmt.Println("Error:",error)
		return
	}

	fmt.Println("Resultado:",result)


}