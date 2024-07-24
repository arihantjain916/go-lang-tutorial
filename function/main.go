package main

import "fmt"

func main() {

	fmt.Println("Welcome to function folder...")
	greetUser("Arihant")

	a := addNum(2, 3)
	fmt.Println("The sum of two number is:", a)

	pro := proAdder(1, 2, 3, 4, 5)
	fmt.Println("The sum of all number is:", pro)

}

func addNum(a int, b int) int {
	return a + b
}

func proAdder(value ...int) int {
	sum := 0
	for _, v := range value {
		sum += v
	}
	return sum
}
func greetUser(name string) {
	fmt.Println("Hello", name)
}
