package main

import "fmt"

func main() {
	// suppose i want to pass a value through a function
	// if i pass directly a copy will be sent
	// to ensure that does not occur we use pointers.

	// var ptr *int
	// fmt.Println("Value of Pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	*ptr = 223
	fmt.Println(myNumber)
}
