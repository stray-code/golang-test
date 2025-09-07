package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums2 = make([]int, 2)
	for i := range 10 {
		nums2 = append(nums2, i)
	}
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	fmt.Println(len(nums2))

	var numbers = make([]int, 0, 5)
	numbers = append(numbers, 2)
	var numbers2 = make([]int, len(numbers))

	copy(numbers2, numbers)

	fmt.Println(numbers, numbers2)
	fmt.Println(cap(numbers), cap(numbers2))

	var array = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(array[3:])

	array1 := []int{1, 2, 3}
	array2 := []int{1, 2, 3, 4}
	fmt.Println(slices.Equal(array1, array2))
}
