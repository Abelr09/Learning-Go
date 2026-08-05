package main

import (
	"fmt"
	"time"
)

/**
chan int      // leer y escribir
<-chan int    // SOLO leer
chan<- int    // SOLO escribir
*/

func worker(id int, tareas <-chan int, resultados chan int) {

	for i := range tareas {
		fmt.Printf("worker %d: procesando tarea %d\n", id, i)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "tarea iniciada", i)
		resultados <- i * 2
	}
}

func main() {

	const numeroDeTareas = 5

	tareas := make(chan int, numeroDeTareas)

	resultados := make(chan int, numeroDeTareas)

	for w := 1; w <= 3; w++ {
		go worker(w, tareas, resultados)
	}

	for j := 1; j <= numeroDeTareas; j++ {
		tareas <- j
	}
	close(tareas)
	for a := 1; a <= numeroDeTareas; a++ {
		<-resultados
	}
}
