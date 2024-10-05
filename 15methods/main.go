package main

import "fmt"

func main() {
	dikshita := User{"Dikshita", "dpatel60@email.com", true, 24}
	dikshita.GetStatus()
	dikshita.NewEmail()
	fmt.Println(dikshita)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active ", u.Status)
}

// here a copy of email is passed actual email is not changed
func (u User) NewEmail() {
	u.Email = "new@email.com"
	fmt.Println("New user email ", u.Email)
}
