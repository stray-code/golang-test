package main

import "fmt"

func changeNum(num int) {
	num = 5

	fmt.Println("in changeNum", num)
}

func changeNumPointer(num *int) {
	*num = 10

	fmt.Println("in changeNumPointer", *num)
}

func main() {
	num := 1

	changeNum(num)

	fmt.Println("after changeNum in main", num)

	numPointer := 2

	changeNumPointer(&numPointer)

	fmt.Println("after changeNumPointer in main", numPointer)
}
