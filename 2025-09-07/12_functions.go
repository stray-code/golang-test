package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"
}

func processIt(fn func(a int)) {
	fn(323)
}

func main() {
	result := add(3, 5)
	fmt.Println(result)

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	fn := func(a int) {
		fmt.Println(a)
	}

	processIt(fn)
}
