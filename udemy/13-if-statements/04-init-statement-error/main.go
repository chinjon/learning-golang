package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// fmt.Prinln(food)
	// will not run "food is scoped inside the if statement"
}
