package main

import "fmt"

func main() {
	fmt.Println("########### if conditional ##########")

	x := 10
	
	if x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is equal to zero")
	} else {
		fmt.Println("x is positive")
	}

	// condition short variable

	if y:=-10; y < 0 {
		fmt.Println("y is negative")
	}

}
