package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	fmt.Println("Welcome to Methods Folder")

	user := User{"Arihant", "arihant@gmail.com", true, 23}
	fmt.Printf("User is %+v\n", user)
	user.getStatus()
	user.setEmail("arihant916@gmail.com")
	fmt.Printf("User is %+v\n", user)
}

func (u User) getStatus() {
	fmt.Println(u.Status)
}

func (u User) setEmail(email string){
	u.Email = email
	fmt.Println(u.Email)

}
