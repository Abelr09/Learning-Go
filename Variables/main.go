package main

import "fmt"

func main() {
	//Declaracion de una variable de tipo string
	var cadena = "Inicial"
	fmt.Println(cadena)
	//Declaracion de una variable de tipo int
	var entero1, entero2 int = 10, 20
	fmt.Println(entero1, entero2)
	//Declaracion de una variable de tipo float
	var flotante float32 = 3.14
	fmt.Println(flotante)
	//Declaracion de una variable de tipo booleano
	var booleano bool = true
	fmt.Println(booleano)

	var enteroSimple int      //Declaracion de una variable de tipo int sin inicializar
	fmt.Println(enteroSimple) //Imprimimos el valor de la variable enteroSimple, que es 0 por defecto

	enteroSimple = 10         //Inicializamos la variable enteroSimple con el valor 10
	fmt.Println(enteroSimple) //Imprimimos el valor de la variable enteroSimple, que ahora es 10

	fruta := "Manzana" //Declaracion e inicializacion de una variable de tipo string usando el operador :=
	fmt.Println(fruta) //Imprimimos el valor de la variable fruta, que es "Manzana"

}
