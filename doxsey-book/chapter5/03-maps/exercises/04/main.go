package main

import "fmt"

// Write a program that finds the smallest number in this list:
// x := []int{
// 	48,96,86,68,
// 	57,82,63,70,
// 	37,34,83,27,
// 	19,97, 9,17,
// }

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	y := x[0]
	for i := 0; i < len(x); i++ {

		if x[i] < y {
			y = x[i]
		}
	}

	fmt.Println(y)
}
