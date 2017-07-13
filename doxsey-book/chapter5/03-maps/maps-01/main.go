package main

import "fmt"

// a map is an unordered collection of key-value pairs
// maps are used to look up a value by its associated key

// var x map[string]int
// the map type is reporesented by the keyword map, followed by the key type in brackets and finally the value type

// "x is a map of strings to ints"

func main() {
	x := make(map[string]int)
	// creates or declares a map with a string-key and a value-integers
	x["key1"] = 10
	// creates an entry with key of "key1" and value of 10

	x["key2"] = 11
	fmt.Println(x)
}
