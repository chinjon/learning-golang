package main

import "fmt"

func main() {
	// the scope for x is narrowed
	// x is limited to this function
	x := 0
	// anonymous function - function without name
	// func expression = assign a function to a variable
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions without closure, for two or more funcs to have access to the same variable, that variable would need to be package scope.
*/
