package main

import "fmt"

func main() {
	fmt.Println("Structs: ")

	// There is no inheritance, super, parent

	dikshita := User{"Dikshita", "dpatel60@email.com", true, 24}
	fmt.Println(dikshita)

	fmt.Printf("dikshita details are Name: %v, Email: %v \n", dikshita.Name, dikshita.Email)

}

// type & struct keyword
// Make capital letter for Name
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
