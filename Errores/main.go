package main

import (
	"errors"
	"fmt"
)

var ErrorDeCafe = fmt.Errorf("No hay cafe")
var ErrorDeEnergia = errors.New("Ya no hay electricidad")

func hacerCafe(args int) error {
	if args == 2 {
		return ErrorDeCafe
	} else if args == 4 {
		return fmt.Errorf("Haciendo Cafe: %w", ErrorDeEnergia)
	}
	return nil
}

func main() {

	for i := range 5 {
		if err := hacerCafe(i); err != nil {

			if errors.Is(err, ErrorDeCafe) {
				fmt.Println("Necesitamos mas cafe")
			} else if errors.Is(err, ErrorDeEnergia) {
				fmt.Println("Ahora no puedo calentar agua")
			} else {
				fmt.Printf("Error desconocido: %s\n", err)
			}
			continue
		}
	}
	fmt.Println("Ya hay cafe!")
}
