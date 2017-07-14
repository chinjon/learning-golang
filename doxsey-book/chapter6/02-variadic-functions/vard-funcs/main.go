package main

import "fmt"

func add(args ...int) int {
	// by using an ellipsis (...) before the type name of the last parameter, you can indicate that it takes zero or more of those parameters.

	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// func main() {
// 	fmt.Println(add(1, 2, 3))
// }

func main() {
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}
