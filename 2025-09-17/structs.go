package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 3}
	v2 = Vertex{X: 2}
	v3 = Vertex{}
	p  = &Vertex{4, 5}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{2, 4}
	v.X = 9
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(p.X)

	fmt.Println(v1, v2, v3, *p)
}
