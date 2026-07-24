package main

//Go no te deja compilar nada si no estas utilizando los paquetes que importas, por eso es importante importar solo lo que vas a usar
import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME") //obtenemos la variable de entorno HOME
	if envVar == "" {
		fmt.Println("HOME environment variable is not set.") //si la variable de entorno no esta seteada, imprimimos un mensaje
	} else {
		fmt.Printf("HOME environment variable: %v\n", envVar) //si la variable de entorno esta seteada, imprimimos su valor
	}

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err) //si hay un error al crear el archivo, imprimimos el error
		return
	}
	defer file.Close() //cerramos el archivo al final de la funcion main
}
