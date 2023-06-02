package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {

	var person1 Person
	var person2 Person

	person1.Name = "Tarcisio"
	person1.Age = 35

	person2.Name = "Fake"
	person2.Age = 20

	fmt.Println("Person 1 Name: ", person1.Name)
	fmt.Println("Person 1 Age: ", person1.Age)
	
	fmt.Println("Person 2 Name: ", person2.Name)
	fmt.Println("Person 2 Age: ", person2.Age)
}