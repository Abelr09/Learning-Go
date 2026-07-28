package main

import "fmt"

func valoresMultiples(nombre1, nombre2, apellido string) (map[string]string, map[string]string) {
	mapa1 := make(map[string]string)
	mapa2 := make(map[string]string)

	mapa1[nombre1] = apellido
	mapa2[nombre2] = apellido

	return mapa1, mapa2
}

func main() {

	nombre1 := "Abel"
	nombre2 := "Antonio"
	apellido := "Amarilla"

	mapaPrimero, mapaSegundo := valoresMultiples(nombre1, nombre2, apellido)

	fmt.Println("Mapa con el primer nombre: ", mapaPrimero)
	fmt.Println("Mapa con el segundo nombre: ", mapaSegundo)

	_, nombre := valoresMultiples(nombre1, nombre2, apellido)
	fmt.Println(nombre)

	concatenar := nombre1 + " " + nombre2 + " " + apellido
	fmt.Println("Nombre completo: ", concatenar)
}
