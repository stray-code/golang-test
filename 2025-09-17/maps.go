package main

import "fmt"

type Vertex struct {
	Lat  float64
	Long float64
}

var m map[string]Vertex

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m == nil)

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m2)
}
