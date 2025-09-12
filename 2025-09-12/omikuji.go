package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	number := rand.N(10)

	switch number {
	case 0:
		fmt.Println("大吉")
	case 1, 2, 3:
		fmt.Println("吉")
	case 4, 5, 6:
		fmt.Println("小吉")
	case 7, 8:
		fmt.Println("凶")
	case 9:
		fmt.Println("大凶")
	}
}
