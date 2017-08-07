package main

import "fmt"

// maps are for key-value association
// like a dictionary
// word = key
// definition = values

// slices for lists
// maps for key-value pairs
// structs for different variables in one container

// map will store a pointer to where the data is stored

func main() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}
