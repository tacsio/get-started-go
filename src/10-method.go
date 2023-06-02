package main

import "fmt"

type rectangle struct {
	width int
	height int
}

//function
func area(r rectangle) int {
	return r.width * r.height
}

//method
func (r rectangle) area() int {
	return r.width * r.height
}

func main() {
	r := rectangle{ width: 10, height: 5 }

	fmt.Println("[Function] The area of the rectangle is: ", area(r))
	fmt.Println("[Method] The area of the rectangle is: ", r.area())
}