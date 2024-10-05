package main

import "fmt"

func main() {
	var fruits [3]string

	fruits[0] = "apple"
	fruits[1] = "mango"

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var veg = [3]string{"potato", "onion", "spinach"}
	fmt.Println(veg)
}
