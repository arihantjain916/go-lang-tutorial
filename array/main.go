package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array Folder...")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[2] = "Mango"
	fruitlist[3] = "Orange"

	fmt.Println("Fruit list is: ", fruitlist)

	var veglist = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("Veg list is: ", veglist)
}
