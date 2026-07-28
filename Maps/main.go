package main

import (
	"fmt"
	"maps"
	"strings"
)

func main() {

	separador := strings.Repeat("-", 20)
	fmt.Println(separador)

	mapa := make(map[string]int)

	mapa["Abel"] = 4
	mapa["Amarilla"] = 8

	fmt.Println(mapa)

	fmt.Println(separador)

	version1 := mapa["Abel"]
	version2 := mapa["Amarilla"]

	fmt.Println(separador)

	fmt.Println("valor: ", version1, version2)
	fmt.Println("Longitud del mapa: ", len(mapa))

	fmt.Println(separador)

	_, dato := mapa["Abel"] // el guion bajo permite ignorar el primer dato obtenido como valor "4" y el dato debe confirmar que afuera de abel exista el dato
	fmt.Println("Dato: ", dato)

	fmt.Println(separador)

	delete(mapa, "Abel") // delete es para eliminar un solo elemento o registro
	fmt.Println("Eliminado Abel: ", mapa)

	fmt.Println(separador)

	clear(mapa)
	fmt.Println("Limpieza total: ", mapa)

	fmt.Println(separador)

	nuevo1 := map[string]int{"Abel": 1, "Amarilla": 2}
	fmt.Println("mapa1: ", nuevo1)

	nuevo2 := map[string]int{"Abel": 1, "Amarilla": 2}
	fmt.Println("mapa2: ", nuevo2)

	if maps.Equal(nuevo1, nuevo2) {
		fmt.Println("Los mapas son iguales")
	} else {
		fmt.Println("Los mapas son diferentes")
	}
}
