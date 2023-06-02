package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {

	person1 := &Person{ Name: "Jhon", Age: 40}
	person2 := &Person{ Name: "Mary", Age: 20}
	person3 := Person{ Name: "Paul", Age: 40}

	fmt.Println(compareAge(person1, person2))
	fmt.Println(compareAge(person1, &person3))


}


func compareAge(p1, p2 *Person) string {
	if p1.Age == p2.Age {
		return "Equal"
	}

	return "Not equal"
}