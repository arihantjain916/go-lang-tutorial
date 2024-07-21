package main

import (
	"fmt"
)

// ** Outside variable declare **
const LoginToken string = "I am Outside Variable"

func main() {
	// *** String ***
	// var username string = "Arihant Jain"
	// fmt.Println(username)
	// fmt.Printf("Variable of is type: %T \n", username)

	// *** Boolean ***
	// var isLogIn bool = true
	// fmt.Println(isLogIn)
	// fmt.Printf("Variable of is type: %T \n", isLogIn)

	// *** Integer ***
	// var num int = 215
	// fmt.Println(num)
	// fmt.Printf("Variable of is type: %T \n", num)

	// *** Float ***
	// var num float32 = 2.158517854785185
	// fmt.Println(num)
	// fmt.Printf("Variable of is type: %T \n", num)

	// *** default values and some aliases ***
	// var defaultInt int
	// var defaultFloat float32
	// var defaultBool bool
	// var defaultString string
	// fmt.Println(defaultInt, defaultFloat, defaultBool, defaultString)

	// *** Implicit Type ***
	// var websiteUrl = "arihantjain916.com"
	// fmt.Println("Website Url is", websiteUrl)
	// fmt.Printf("Variable of is type: %T \n", websiteUrl)

	// *** No var style ***
	// numberOfUser := "300000"
	// fmt.Println(numberOfUser)

	// *** Accessing Outside Token ***
	fmt.Println(LoginToken)
}
