package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointer's Folder...")

	myNum := 23
	var ptr = &myNum
	fmt.Println(*ptr)
	fmt.Println(ptr)

	*ptr = *ptr + 5
	fmt.Println(myNum)
}
