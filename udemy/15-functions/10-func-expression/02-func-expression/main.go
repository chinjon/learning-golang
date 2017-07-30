package main

import "fmt"

func main() {

	// greeting variable being assigned an anonymous function
	// called a func expression

	greeting := func() {
		fmt.Println("Hello World")
	}

	greeting()
}
