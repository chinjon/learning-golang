package main

import "fmt"

func main() {
	var smallNum, largeNum int

	fmt.Print("Enter a small number")
	fmt.Scan(&smallNum)
	fmt.Print("Enter a large number")
	fmt.Scan(&largeNum)
	fmt.Printf("The remainder for %v and %v is %v \n", largeNum, smallNum, largeNum/smallNum)
}
