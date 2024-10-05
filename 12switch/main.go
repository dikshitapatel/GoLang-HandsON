package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randomDice := rand.Intn(6) + 1

	fmt.Println(randomDice)

	switch randomDice {
	case 1:
		fmt.Println("Dice 1")
	case 6:
		fmt.Println("Game Starts")
		fallthrough //runs the subsequent cases as well
	default:
		fmt.Println("Dice value ", randomDice)

	}

}
