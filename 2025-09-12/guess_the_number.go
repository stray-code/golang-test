package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("guess the number game started")
	fmt.Println("guess the secret number (form 0 to 99) to win")
	answer := rand.N(100)

	for range 7 {
		fmt.Println("input the answer")

		var input int

		fmt.Scan(&input)

		if input == answer {
			fmt.Println("correct!!")
			break
		}

		if input > answer {
			fmt.Println("higher")
			continue
		}

		if input < answer {
			fmt.Println("lower")
		}
	}

	fmt.Printf("the answer was %d.", answer)
}
