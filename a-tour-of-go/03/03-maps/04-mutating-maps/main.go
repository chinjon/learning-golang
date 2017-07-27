package main

import "fmt"

// insert or update an element in map m:
// m[key] = elem

// retrieve an element:
// elem = m[key]

// delete an element:
// delete(m, key)

// test that a key is present with a two-value assignment:
// elem, ok = m[key]

// if key is in m, ok is true
// if not, ok is false

// if key is not in map, then elem is the zero value for the map's element type

// elem, ok := m[key]

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
