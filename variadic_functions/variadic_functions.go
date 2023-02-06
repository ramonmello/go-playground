package main

import "fmt"

func variadicFunc(arg ...int) {
	for i := range arg {
		fmt.Println(arg[i])
	}
}

func printAll(arg ...interface{}) {
	fmt.Println(arg...)
}

func v1(names ...string) { // pack operator used before type in argument
	fmt.Println(names)
}

func main() {
	// Variadic functions exemple
	variadicFunc(1, 2, 3, 4)
	printAll(1, "John", 0.95, 'a')

	// Pack and unpack operator
	var names []string = []string{"Albert", "Issac"}

	v1("John", "Jane", "Dexter", "Bruce") // [John Jane Dexter Bruce]

	// Here is unpack operator in action
	v1(names...) // [Albert Issac]
}
