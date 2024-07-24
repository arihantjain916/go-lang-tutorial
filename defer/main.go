package main

import "fmt"

func main() {

	fmt.Println("Welcome to Defer Folder...")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	fmt.Println("Four")
	fmt.Println("Five")
	fmt.Println("Six")

	myDefer()


}

func myDefer() {
	count := 10
	for i := 1; i <= count; i++ {
		defer fmt.Println(i)
	}
}

// Defer make the code line or func run in last and in reverse order
