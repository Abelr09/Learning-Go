package main

import "fmt"

func main() {
	var numero int = 10
	numero2 := 20
	fmt.Println(numero, numero2)

	var numeroEntero = 10
	var numeroDoble = 20.5

	resultado := numeroEntero + int(numeroDoble)    // Convertimos numeroDoble a int antes de la suma
	fmt.Println("Resultado de la suma:", resultado) // Imprimimos el resultado de la suma la cual es 30 ya que se convierte el valor de numeroDoble a 20 al ser convertido a int

	var nombre string = "Juan"
	apellido := "Pérez"
	nombreCompleto := nombre + " " + apellido
	fmt.Println("Nombre completo:", nombreCompleto) // Imprimimos el nombre completo concatenando nombre y apellido

	var miNombre string = "Abel"
	var miEdad int = 24
	resultado2 := fmt.Sprintf("%s", "Mi nombre es: "+miNombre+" y mi edad es: "+fmt.Sprint(miEdad)) // Convertimos miEdad a string antes de la concatenación
	fmt.Println("Resultado 2:", resultado2)
}
