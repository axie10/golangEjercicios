package greetings

import (
	"fmt"
)

//hello devuelve un saludo para la persona especifica

func Hello (name string)string{
	message := fmt.Sprintf("Hola, %v, bienvenido", name)
	return message
}