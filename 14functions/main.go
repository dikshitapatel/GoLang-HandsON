package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in GOlang")
	greet()

	result := add(3, 5)

	fmt.Println("Result is ", result)

	result = add2(3, 5, 2)

	fmt.Println("Result is ", result)
	var msg string
	result, msg = add3(3, 5, 2)

	fmt.Println("Result is ", result)
	fmt.Println("Message is ", msg)

}

// functions inside functions not allowed
func greet() {
	fmt.Println("Namaste in GOlang")
}

func add(a int, b int) int {
	return a + b
}

// Many parameters
func add2(values ...int) int {
	total := 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	return total
}

// Return mukltiple values
func add3(values ...int) (int, string) {
	total := 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	return total, "Addition"
}
