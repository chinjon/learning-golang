package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// range form of the for loop iterates over a slice or map

// when ranging over a slice, two values are returned for each iteration.
// the first is the index
// the second is a copy of the element at the index

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
