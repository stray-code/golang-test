package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	sum := 0

	for _, num := range nums {
		fmt.Printf("num: %d\n", num)
		sum += num
	}

	fmt.Printf("sum: %d\n", sum)

	object := map[string]string{"fname": "john", "lname": "doe"}

	for key, value := range object {
		fmt.Println(key, value)
	}

	for index, code := range "golang" {
		fmt.Println(index, code, string(code))
	}
}
