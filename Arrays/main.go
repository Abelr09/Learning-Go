package main

import "fmt"

func main() {

	var arreglo [5]int
	fmt.Println("Arreglo completo:", arreglo)

	arreglo[4] = 100
	fmt.Println("Arreglo con valor modificado:", arreglo)

	fmt.Println("Elemento en la posición 4 es igual a:", len(arreglo))

	lista := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lista completa:", lista)

	unlimit := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Lista sin limite:", unlimit)
	fmt.Println("Tamanio del arreglo inferido: ", len(unlimit))

}
