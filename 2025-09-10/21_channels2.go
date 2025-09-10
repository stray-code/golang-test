package main

import (
	"fmt"
)

func sum(result chan int, num1 int, num2 int) {
	result <- num1 + num2
}

func main() {
	result := make(chan int)

	go sum(result, 3, 4)

	res := <-result
	fmt.Println(res)
}
