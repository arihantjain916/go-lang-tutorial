package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Tutorial")

	// Time
	tiime := time.Now()
	fmt.Println(tiime.Format("01-02-2006 Monday"))

	// Note: Read the docs...
}
