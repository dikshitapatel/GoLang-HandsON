package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Friday"}
	fmt.Println(days)

	fmt.Println("*************")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("*************")
	for i := range days { //range gives index
		fmt.Println(days[i])
	}

	fmt.Println("*************")
	for key, value := range days {

		fmt.Printf("Key is: %v , value is %v \n", key, value)

	}

	fmt.Println("*************")
	start := 1

	for start < 4 {

		if start == 2 {
			goto label
		}
		fmt.Println(days[start])
		start++
	}

label:
	fmt.Println("2 reached")
	//++start not allowed in GO

}
