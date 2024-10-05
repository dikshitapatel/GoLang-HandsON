package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio package

func main() {
	//Taking input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for your pizza")

	// comma ok syntax
	// _ is the error
	input, _ := reader.ReadString('\n')

	fmt.Println("Rating", input)
}
