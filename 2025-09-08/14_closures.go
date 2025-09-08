package main

import "fmt"

func counter() func() int {
	var count int

	return func() int {
		count++

		return count
	}
}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
