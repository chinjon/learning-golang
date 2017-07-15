// Write a function with one variadic parameter that finds the greatest number in a list of numbers.

package main

import "fmt"

func sortLargest(input ...int) int {
	largest := 0
	for _, i := range input {
		if i > largest {
			largest = i
		}
	}

	return largest
}

func main() {
	fmt.Println(sortLargest(1, 12, 6, 8, 21))
	fmt.Println(sortLargest(3, 43, 44, 76, 77, 1, 2, 3, 4))
}
