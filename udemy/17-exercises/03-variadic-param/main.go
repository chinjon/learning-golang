package main

import "fmt"

func findGreatest(nums ...int) int {
	var greatest int

	for _, n := range nums {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}

func main() {
	x := findGreatest(1, 10, 32, 2, 4, 3)
	fmt.Println(x)
}
