package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices:")
	var fruits = []string{"cherry"}

	fruits = append(fruits, "apple", "mango")

	fmt.Println(fruits)
	fruits = append(fruits[0:2])
	fmt.Println(fruits)

	//make keyword to create append

	highscores := make([]int, 4)

	highscores[0] = 100
	highscores[1] = 200
	highscores[2] = 300
	highscores[3] = 400

	fmt.Println(highscores)

	highscores = append(highscores, 444, 55, 66)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// how to remove value from slice based on index

	var courses []string
	courses = append(courses, "Python", "Java", "Go", "JS")
	fmt.Println(courses)

	var index = 2
	courses = append(courses[:index], courses[index+1:]...) // This removes Go
	fmt.Println(courses)

}
