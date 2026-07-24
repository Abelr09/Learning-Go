package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	separador := strings.Repeat("-", 20) //strings.Repeat es una funcion que repite un string n veces, en este caso 20 veces

	fmt.Println(separador)

	fmt.Println("Escribe un numero del 1 al 3")
	i := 2
	fmt.Println("Escribe", i, "como")

	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	}

	fmt.Println(separador)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Es dia laboral")
	}

	fmt.Println(separador)

	tiempo := time.Now()
	fmt.Println("La hora es: ", tiempo)
	switch {
	case tiempo.Hour() < 12:
		fmt.Println("Es de mañana")
	default:
		fmt.Println("Es de tarde")
	}
}
