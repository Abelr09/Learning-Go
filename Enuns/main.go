package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var StateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (estadoServidor ServerState) String() string {
	return StateName[estadoServidor]
}

func main() {
	redServidor := verificacionDeRed(StateIdle)
	fmt.Println("Estado del servidor: ", redServidor)

	segundaRevision := verificacionDeRed(redServidor)
	fmt.Println("Segunda revision: ", segundaRevision)
}

func verificacionDeRed(servidor ServerState) ServerState {

	switch servidor {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Estado desconocido: %s", servidor))
	}
}
