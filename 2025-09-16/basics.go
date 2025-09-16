package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

var c, python bool
var java bool = true

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("my favorite number is ", rand.IntN(10))
	fmt.Printf("now you have %g probrems\n", math.Sqrt(2))
	fmt.Println("PI", math.Pi)
	fmt.Println("23 + 53 =", add(23, 53))
	a, b := swap(0, 10)
	fmt.Println("swap", a, b)
	var i int
	fmt.Println(i, c, python, java)
	f := math.Sqrt(float64((1 * 1) + (4 * 4)))
	z := uint(f)
	fmt.Println(f, z)
	v := 34
	fmt.Printf("v is of type %T\n", v)
	const world = "世界"
	fmt.Println("Hello", world)
}
