package main

import "fmt"

func main() {
	// Declare an integer variable
	num := 5

	// Declare a pointer to the integer variable
	var numPointer *int

	// Assign the memory address of the integer variable to the pointer
	numPointer = &num

	// Print the value of the integer variable using the pointer
	fmt.Println("Value of num using numPointer:", *numPointer)

	// Change the value of the integer variable using the pointer
	*numPointer = 10

	// Print the new value of the integer variable
	fmt.Println("New value of num:", num)
}
