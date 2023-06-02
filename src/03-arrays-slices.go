package main

import (
	"fmt"
	"reflect"
)

func main() {

	//quando adicionamento o tamanho => array
	arrayGo := [3]string { "apple", "banana", "cherry"}
	fmt.Println(arrayGo)


	//slices: Quando não informamos tamanho
	sliceGo := []string { "blue", "red", "green"}
	sliceGo = append(sliceGo, "orange", "white")

	fmt.Println(sliceGo)

	fmt.Println("")

	slice_compare()
}

func slice_compare() {
	a := []int {1,2,3}
	b := []int {1,2,3}
	c := []int {1,2,4}

	fmt.Println(reflect.DeepEqual(a,b)) //true	
	fmt.Println(reflect.DeepEqual(b,c)) //false
}