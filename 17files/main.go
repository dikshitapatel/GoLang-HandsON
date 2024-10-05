package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This text goes into a file Hi Hello World!"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
		//Shut downs execution and shows error
	}
	length, err := io.WriteString(file, content)

	checkErr(err)

	fmt.Println("Length is", length)

	defer file.Close()

	readFile("./myfile.txt")

}

func readFile(fileName string) {

	content, err := os.ReadFile(fileName)
	checkErr(err)
	fmt.Println("File content is", string(content))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
		//Shut downs execution and shows error
	}
}
