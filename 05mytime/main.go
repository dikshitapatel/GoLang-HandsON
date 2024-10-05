package main

import (
	"fmt"
	"time"
)

// time package
func main() {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// The format is fixed 01-02-2006 15:04:05 Monday :( annoying
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

}
