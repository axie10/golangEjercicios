package main

import (
	"fmt"
	"log"

	"github.com/axie10/greeting"
)

func main() {

	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"Alex", "Roel", "Juan"}
	messages, err := greeting.Hellos(names)
	if err != nil {
		log.Fatal()
	}

	// message, err := greeting.Hello("Manu")
	// if err != nil {
	// 	log.Fatal()
	// }

	fmt.Println(messages)

}
