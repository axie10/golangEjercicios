package main

import (
	// "errors"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	// "fmt"
	// "strconv"
)

//estructura de contactos
type Contact struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//Guardar contactos en un archivo
func saveContactsToFile(contacts []Contact) error{

	file, err := os.Create("contact.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil{
		return err
	}

	return nil
}

//funcion para cargar los contactos del archivo
func loadContactFromFile(contacts *[]Contact) error{

	file, err := os.Open("contact.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}


func main(){

	//slide de contactos
	var contacts []Contact

	//cargamso los contactos del file
	err := loadContactFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	//para ingresar los contactos creamos una instancia de bufio
	reader := bufio.NewReader(os.Stdin)

	for {
		//variable opciones
		var opcion int
		//menu
		fmt.Print("===Gestor de contactos===\n",
			"1. Agregar contactos\n",
			"2. Mostrar contactos\n",
			"3. Salir\n",
			"Elige una opcion: ")
		
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		//Manejar las opciones del usuario
		switch opcion {
		case 1:
			//ingresamos el contacto
			var c Contact
			fmt.Println("Nombre: ")
			c.Name,_ = reader.ReadString('\n')
			fmt.Println("Email: ")
			c.Email,_ = reader.ReadString('\n')
			fmt.Println("Phone: ")
			c.Phone,_ = reader.ReadString('\n')

			//agregamos el contacto
			contacts = append(contacts, c)

			//guardar contacto
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}

		case 2:
			//Mostrar todos los datos del contacto
			fmt.Println("=======================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
				index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=======================")
		case 3:
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
}