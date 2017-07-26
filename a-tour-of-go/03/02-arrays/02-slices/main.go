package main

import "fmt"

// an array has a fixed size.
// a slice is a dynamically-sized, flexible view into the elements of an array

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[1:4]
	fmt.Println(s)
}
