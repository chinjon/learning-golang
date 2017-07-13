package main

import (
	"fmt"
)

func main() {
	// var x []float64

	// x := make([]float64, 5)
	// creates a slice with an underlying float64 array of length 5

	x := make([]float64, 5, 10)
	// a slice of length 5 with a capacity of 10

	fmt.Println(x)
}
