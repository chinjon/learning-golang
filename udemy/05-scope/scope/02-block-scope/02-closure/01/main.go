package main

import "fmt"

func main() {
	x := 42

	fmt.Println(x)
	{
		fmt.Println(x) // will work, x is scoped outside
		y := "This is a string"
		fmt.Println(y)
	}

	// fmt.Println(y) // outside scope of y
}
