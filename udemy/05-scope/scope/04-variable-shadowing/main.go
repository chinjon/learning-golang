package main

import "fmt"

// don't do this, this is bad coding practice
// there should be a uniqure name for every variable

func max(x int) int {
	// the int parameter is only available to the body of the function
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result, not the function
}
