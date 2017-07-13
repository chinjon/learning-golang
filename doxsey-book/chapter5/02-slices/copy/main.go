package main

import "fmt"

// slices are typically used to reporesent lists of items, particularly when you need to access the nth item quickly

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) // creates array [0,0]
	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)

	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	// copy takes two arguments dst and src
	// all of the entries in src are copied into dst overwritting whatever is there.

	// if the lengths of the two slices are not the same, the small of the two will be used
}
