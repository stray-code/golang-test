package main

import "fmt"

func main() {
	fmt.Println("数字を入力してください")
	var input int
	fmt.Scan(&input)
	fmt.Printf("入力は%d", input)
}
