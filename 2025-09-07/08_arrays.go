package main

import "fmt"

func main() {
	var nums [4]int

	fmt.Println(len(nums))

	nums[0] = 1

	fmt.Println(nums)

	var values [3]bool
	values[2] = true
	fmt.Println(values)

	nums2d := [2][2]int{{5, 3}, {4, 2}}
	fmt.Println(nums2d)
}
