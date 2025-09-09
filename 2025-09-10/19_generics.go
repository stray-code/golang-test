package main

import "fmt"

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T comparable] struct {
	elements []T
}

func main() {
	// nums := []int{1, 2, 3}
	// names := []string{"golang", "javascript"}
	values := []bool{true, false, false}

	printSlice(values)

	myStack := stack[string]{
		elements: []string{"golang", "java"},
	}

	fmt.Println(myStack)
}
