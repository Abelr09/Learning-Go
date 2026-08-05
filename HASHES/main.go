package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	cadenaBase := "Abel amarilla"

	hash := sha256.New()
	hash.Write([]byte(cadenaBase))

	arregloHash := hash.Sum(nil)
	fmt.Println(cadenaBase)
	fmt.Printf("%x\n", arregloHash)

	pass := "1234"
	passHash := hash.Sum(nil)
	fmt.Println(pass)
	fmt.Printf("%x\n", passHash)
}
