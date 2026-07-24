package main

import "fmt"

func main() {

	nombre := "Abel"
	edad := 24

	if nombre == "Abel" {
		fmt.Println("Hola Abel")
	} else { //el else siempre debe ir en la misma línea que la llave de cierre del if, de lo contrario se genera un error de sintaxis
		fmt.Println("Hola desconocido")
	}
	if edad >= 18 {
		fmt.Println("Eres mayor de edad, puedes votar")
	} else {
		fmt.Println("Eres menor de edad, no puedes votar")
	}
	if edad >= 18 && nombre == "Abel" {
		fmt.Println("Bienvenido Abel")
	} else {
		fmt.Println("Usuario desconocido")
	}

	if numero := 88; numero > 9 { //Go permite declarar variables dentro de la condición del if, estas variables solo existen dentro del bloque del if
		fmt.Println(numero, "El número es de dos digitos")
	} else if numero < 10 {
		fmt.Println(numero, "El número es de un solo digito")
	} else {
		fmt.Println(numero, "El número es igual a 9")
	}

}
