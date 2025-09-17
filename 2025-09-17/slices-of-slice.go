package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][0] = "O"
	board[2][0] = "X"
	board[0][1] = "O"
	board[1][1] = "X"
	board[2][1] = "o"

	for _, b := range board {
		fmt.Printf("%s\n", strings.Join(b, " "))
	}
}
