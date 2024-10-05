package main

import "fmt"

func main() {

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["J"] = "Java"
	languages["CSS"] = "Cascading Style Sheets"

	fmt.Println(languages)

	//Get value from key
	fmt.Println(languages["JS"])

	delete(languages, "PY")
	fmt.Println(languages)

	//Printing key values in loop
	for key, value := range languages {

		fmt.Printf("Key is: %v , value is %v \n", key, value)

	}
}
