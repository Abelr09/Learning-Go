package main

import "fmt"

func main() {
	mensajes := make(chan string)
	go func() {
		mensajes <- "hola"
	}()

	mensaje := <-mensajes
	fmt.Println(mensaje)

	go func() { mensajes <- "pong" }()

	segundoMensaje := <-mensajes
	fmt.Println(segundoMensaje)
}
