package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slice Folder...")

	// var fruitList = []string{}

	// fruitList = append(fruitList, "Apple", "Tomato", "Peach")
	// fmt.Println("FruitList is ", fruitList)

	// To remove any value from array
	// fruitList = append(fruitList[1:])
	// fmt.Println("FruitList is ", fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	highScore = append(highScore, 555, 666, 777)

	fmt.Println("HighScore is ", highScore)

	sort.Ints(highScore)
	fmt.Println("Sorted HighScore is ", highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
}
