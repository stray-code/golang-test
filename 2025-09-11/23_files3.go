package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}

	for _, f := range fileInfo {
		fmt.Println(f.Name(), f.IsDir())
	}
}
