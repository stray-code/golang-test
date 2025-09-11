package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("dist.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted successfully")
}
