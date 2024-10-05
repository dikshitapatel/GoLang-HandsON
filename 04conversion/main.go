package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for your pizza 1 to 5")

	// comma ok syntax
	// _ is the error
	input, _ := reader.ReadString('\n')
	//Parsing string to flot and trim space to remove /n
	fmt.Println("Thanks for Rating ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}
}
