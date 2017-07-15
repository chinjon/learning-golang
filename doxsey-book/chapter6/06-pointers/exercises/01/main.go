// sum is a function that takes a slice of numbers and adds them together. What would its function signature look like in Go?

package main

import "fmt"

func sum(input ...int) int {
	sum := 0

	for _, i := range input {
		sum += i
	}

	fmt.Println(sum)
	return sum
}

func main() {
	sum(1, 3, 4, 4, 5) // result is 17
}
