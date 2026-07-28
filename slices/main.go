package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	/*
		Los slices sirven para
	*/

	separador := strings.Repeat("-", 20)
	fmt.Println(separador)

	var arregloCadenas []string
	fmt.Println("datos: ", "contenido: ", arregloCadenas, "condicion: ", arregloCadenas == nil, "longitud: ", len(arregloCadenas) == 0)

	fmt.Println(separador)

	arregloCadenas = make([]string, 5)
	arregloCadenas[0] = "Hola"
	arregloCadenas[1] = "Mundo"
	fmt.Println("datos: ", arregloCadenas)
	fmt.Println("dato especifico: ", arregloCadenas[2])
	fmt.Println("longitud: ", len(arregloCadenas))
	fmt.Println(separador)
	arregloCadenas = append(arregloCadenas, "f")
	arregloCadenas = append(arregloCadenas, "g")
	fmt.Println("longitud:", len(arregloCadenas))
	fmt.Println(separador)
	sefundoArreglo := []string{"h", "i", "j"}
	fmt.Println("datos 2", sefundoArreglo)

	tercerArreglo := []string{"h", "i", "j"}
	if slices.Equal(sefundoArreglo, tercerArreglo) {
		fmt.Println("2 es igual que 3")
	}

	fmt.Println(separador)

	for i := len(arregloCadenas); i < 7; i++ {
		arregloCadenas = append(arregloCadenas, sefundoArreglo...)
		fmt.Println("Agregando las demas letras: ", arregloCadenas)
		fmt.Println("El total final es: ", len(arregloCadenas))
	}

}
