package main

func main() {
	// Declare an integer variable
	num := 5

	// Declare a pointer to the integer variable
	var numPointer *int

	// Assign the memory address of the integer variable to the pointer
	numPointer = &num

	// Print the value of the integer variable
	println("Value of num:", num)

	// Print the memory address of the integer variable
	println("Memory address of num:", &num)

	// Print the value stored in the pointer (which is the memory address of the integer variable)
	println("Value stored in numPointer:", numPointer)

	// Print the value of the integer variable using the pointer
	println("Value of num using numPointer:", *numPointer)
}
