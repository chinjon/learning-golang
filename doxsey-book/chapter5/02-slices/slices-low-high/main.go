package main

import "fmt"

func main() {
	arr := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// x := arr[0:6]
	// [1,2,3,4,5,6]

	x := arr[1:len(arr)]
	// [2,3,4,4,5,6,7,8,9,10]

	fmt.Println(x)
}
