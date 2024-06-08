package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}

func jugar() {

	numeroAleatorio := rand.Intn(10) + 1
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {

		fmt.Printf("Ingrese un numero: ")
		fmt.Scanln(&numIngresado)

		intentos++
		fmt.Printf("El numero de intentos que le quedan es: %d\n", maxIntentos-intentos)

		if numeroAleatorio == numIngresado {
			fmt.Println("Ha acertado el numero")
			jugarNuevamente()
			return
		} else if numIngresado < numeroAleatorio {
			fmt.Println("El numero que buscamos es mayor que el ingresado")
		} else if numIngresado > numeroAleatorio {
			fmt.Println("El numero que buscamos es menor que el ingresado")
		}
	}

	fmt.Printf("Ha terminado con el numero de intentos el numero que buscamso es: %d", numeroAleatorio)
}

func jugarNuevamente() {

	var eleccion string
	fmt.Printf("Desea jugar otra vez? (s/n):")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar")
	}
}
