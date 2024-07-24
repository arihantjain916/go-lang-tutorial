package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Status bool
}

func main() {

	fmt.Println("Welcome to Struct Folder")

	arihant := User{"Arihant", 22, false}
	fmt.Println(arihant)
	fmt.Printf("Arihant details are: %+v\n", arihant)
	fmt.Printf("Name is: %v and age is: %v\n", arihant.Name, arihant.Age)
}
