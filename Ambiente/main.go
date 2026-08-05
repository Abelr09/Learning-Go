package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error cargando el archivo.env")
	}
	appName := os.Getenv("APP_NAME")
	appEnv := os.Getenv("APP_ENV")
	port := os.Getenv("APP_PORT")

	fmt.Println("Nombre de la aplicacion: ", appName)
	fmt.Println("Entorno de la aplicacion: ", appEnv)
	fmt.Println("Puerto de la aplicacion: ", port)

}
