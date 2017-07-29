package main

import "fmt"

func main() {
	if false {
		fmt.Println("First print statement")
	} else if false {
		fmt.Println("Second print statement")
	} else if true {
		fmt.Println("print statement")
	} else {
		fmt.Println("Third print statement")
	}
}
