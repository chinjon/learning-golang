package main

import "fmt"

func main() {
	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83

	x := [5]float64{98, 93, 77, 82, 83}

	var total float64

	// range keyword is followed by the name of the variable we want to loop over.
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))
	// result === 86.6
}
