package main

import "fmt"

func main() {

	// result, err := divide(10, 0)
	result, err := divide(10, 5)

	if err != nil {
		panic(err)
	}
	fmt.Println("Result: " , result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0 , fmt.Errorf("cannot divide by 0")
	}

	return a/b, nil
}