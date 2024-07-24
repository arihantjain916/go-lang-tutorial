package main

import "fmt"

func main() {

	fmt.Println("Welcome to Loops Folder...")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// 	for _, value := range days {
	// 		fmt.Println(value)
	// }

	a := 1
	for a <= 10 {

		if a == 5 {
			goto lco
		}
		
		if a == 8 {
			break
		}

		if a == 2 {
			a++
			continue
		}

		fmt.Println(a)
		a++
	}

lco:
	fmt.Println("Hello World")
}
