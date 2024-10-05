package main

import "fmt"

func main() {
	var x = 3
	var ptr *int = &x
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println(x)

}
