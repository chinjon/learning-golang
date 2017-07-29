package main

import "fmt"

// Variadic Parameter
// the final parameter in a function signature may have a type prefixed with ...
// a function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
