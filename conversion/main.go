package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Conversion...")

	// Taking user Input
	fmt.Print("Enter your number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// trim space to eliminate any space after input
	input = strings.TrimSpace(input)

	// Convert Input to float

	cinput, err := strconv.ParseFloat(input, 32)
	fmt.Println("Your number is", cinput)
	fmt.Printf("Type of input is %f and type is: %T\n", cinput, cinput)

	if err != nil {
		fmt.Println("error: ", err)
	}
}
