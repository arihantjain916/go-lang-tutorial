package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Map Folder...")

	// Initialize map
	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages: ", language["JS"])

	// Delete Values in Map
	delete(language, "JS")
	fmt.Println("List of all languages: ", language)

	// Print the key and value
	for key, value := range language {
		fmt.Printf("For key %v, the value is: %v \n", key, value)
	}
}
