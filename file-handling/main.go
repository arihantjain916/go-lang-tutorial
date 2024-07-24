package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to File Handling Folder...")

	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, "Arihant Jain")

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./test.txt")
}

func readFile(filepath string) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	fmt.Println("File Content:", string(data))
}
