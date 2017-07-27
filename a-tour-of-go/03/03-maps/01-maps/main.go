package main

import "fmt"

// a map maps keys to values

// the zero value of a map is nil.
// a nil map has no keys, nor can keys be added

// Vertex struct
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
