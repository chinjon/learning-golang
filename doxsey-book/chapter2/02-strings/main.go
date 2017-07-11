package main

import "fmt"

func main() {
	phrase := "Hello, World"
	fmt.Println(len(phrase))
	fmt.Println(phrase[1])         // prints 101 instead of "e"
	fmt.Println(string(phrase[1])) // prints "e"
	fmt.Println("Hello, " + "World")
}
