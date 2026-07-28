package main

import "fmt"

func suma(a int, b int) int {
	resultado := a + b
	return resultado
}

func sumaL(a, b, c int) int {
	return a + b + c

}

func main() {

	var a, b, c int

	fmt.Println("Ingresa el primer numero: ")
	fmt.Scanln(&a) //El metodo scanln es para qye el usuario pueda ingresar los datos
	fmt.Println("Ingresa el segundo numero: ")
	fmt.Scanln(&b)
	fmt.Println("Ingresa el tercer numero: ")
	fmt.Scanln(&c)

	resultado := suma(a, b)
	fmt.Println("La suma de los 2 numeros es: ", resultado)

	resultadoL := sumaL(a, b, c)
	fmt.Println("La suma de los 2 numeros es: ", resultadoL)
}
