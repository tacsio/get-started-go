package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/equality"
)

type Person struct {
	Name string
	Age int
}

func main() {

	person1 := Person{ Name: "jhon", Age: 40 }
	person2 := Person{ Name: "jhon", Age: 40 }
	person3 := Person{ Name: "jhon", Age: 35 }

	fmt.Println(equality.Semantic.DeepDerivate(person1, person2)) //true
	fmt.Println(equality.Semantic.DeepDerivate(person1, person3)) //false
}