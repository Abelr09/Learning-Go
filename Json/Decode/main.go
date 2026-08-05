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
	jsonString := `{"name":"Abel Amarilla","age":24,"email":"abelamaarillla51777@gmail.com"}`

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Nombre: ", person.Name)
	fmt.Println("Edad: ", person.Age)
	fmt.Println("Email: ", person.Email)
}
