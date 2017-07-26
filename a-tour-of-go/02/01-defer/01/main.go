package main

import "fmt"

func main() {
	defer fmt.Println("This goes last?")
	defer fmt.Println("world")
	defer fmt.Println("Hello")
	defer fmt.Println("this defer goes first")
}
