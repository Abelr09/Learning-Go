package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person{

		Name:  "Abel Amarilla",
		Age:   24,
		Email: "abelamaarillla51777@gmail.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return
	}
	fmt.Println("Tu archivo JSON es: ", string(jsonData))
}
