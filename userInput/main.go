package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Taking Input from User")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	// comma ok || err ok
	name, _ := reader.ReadString('\n')
	fmt.Print("Hello", name)
}
