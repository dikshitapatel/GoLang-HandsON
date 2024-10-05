package main

import "fmt"

func main() {
	//Prints Hello Two One World
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

}
