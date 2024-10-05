package main

import "fmt"

func main() {

	loginCount := 12
	var result string
	if loginCount < 0 {
		result = "Regaular user"
	} else {
		result = "Frequent user"
	}
	fmt.Println(result)
}
