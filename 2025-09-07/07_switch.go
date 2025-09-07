package main

import (
	"fmt"
	"time"
)

func main() {
	index := 1

	switch index {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI(3.4)
}
