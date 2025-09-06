package main

import "fmt"

func main() {
	i := 0

	for i < 3 {
		fmt.Printf("i: %d\n", i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Printf("j: %d\n", j)
	}

	for k := 0; k < 3; k++ {
		if k == 1 {
			// break
			continue
		}

		fmt.Printf("k: %d\n", k)
	}

	for l := range 2 {
		fmt.Printf("l: %d\n", l)
	}
}
