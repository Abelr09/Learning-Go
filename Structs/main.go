package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func nuevaPersona(nombre string, edad int) *persona {
	nuevoIndividuo := persona{nombre: nombre}
	nuevoIndividuo.edad = 23
	return &nuevoIndividuo
}

func main() {
	fmt.Println(persona{"Abel", 43})

	fmt.Println(persona{nombre: "Abel", edad: 43})

	fmt.Println(persona{nombre: "Abel"})

	fmt.Println(persona{"River", 50})

	personita := persona{nombre: "DANU", edad: 12}
	fmt.Println(personita.nombre)

	edadPersonita := &personita
	fmt.Println(edadPersonita.edad)

	personita.edad = 8
	fmt.Println(edadPersonita.edad)
	fmt.Println(personita)

}
