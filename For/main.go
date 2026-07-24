package main

import (
	"fmt"
	"strings"
)

func main() {

	separador := strings.Repeat("-", 20) //strings.Repeat es una funcion que repite un string n veces, en este caso 20 veces
	fmt.Println(separador)

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++ // puedes usar i += 1 o i = i + 1 para incrementar el valor de i
	} //el for en Go no tiene while, pero se puede simular con un for y una condición

	fmt.Println(separador)

	for numero := 0; numero <= 5; numero++ { //el for en Go tiene la sintaxis de for(inicializacion; condicion; incremento)
		fmt.Println(numero)
	}

	fmt.Println(separador)

	for rango := 0; rango <= 10; rango += 2 {
		fmt.Println("rango: ", rango)
	}

	fmt.Println(separador)

	for {
		fmt.Println("loop")
		break //el break rompe el loop, si no se pone el loop seria infinito
	}

	fmt.Println(separador)

	for valor := range 6 {
		if valor%2 == 0 {
			continue //el continue salta a la siguiente iteracion del loop
		}
		fmt.Println("valor impar: ", valor)
	}
}
