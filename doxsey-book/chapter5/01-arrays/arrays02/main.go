package main

import "fmt"

func main() {

	// first we create an array and declare a length of 5 to hold numbers
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	// we declare the type of the "total" variable
	var total float64

	// we setup the loop to compute the score
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}
