package main

/*
Los punteros en Go son variables que almacenan la dirección de memoria de otra variable.
En lugar de trabajar directamente con el valor de una variable, los punteros permiten manipular su contenido a través de su dirección de memoria.
Esto es útil para optimizar el rendimiento,
especialmente en sistemas con alta concurrencia o cuando se trabaja con hardware.
*/
import "fmt"

func modificarArreglo(arg *[5]int) {
	(*arg)[0] = 42
}

func main() {
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numeros: ", numeros)

	modificarArreglo(&numeros)
	fmt.Println(numeros)

	numeros[1] = 10
	fmt.Println(numeros)
}
