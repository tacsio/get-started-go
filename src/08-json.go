package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string `json: "PersonName"`
	Age int		`json: "PersonAge"`
}

func main() {
	p1 := Person{ Name: "Jhon", Age: 20}
	json, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Json Data:")
	fmt.Println(string(json))
}