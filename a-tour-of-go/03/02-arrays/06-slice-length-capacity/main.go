package main

import "fmt"

// a slice has both a length and a capacity

// the capacity of a slice is the number of elements in the underlying array, counting from the first element of the slice.

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	b := s[:0]
	printSlice(b)

	c := s[:4]
	printSlice(c)

	d := s[2:]
	printSlice(d)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
