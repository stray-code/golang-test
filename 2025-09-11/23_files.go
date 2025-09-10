package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())

	buf := make([]byte, 12)
	data, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for _, b := range buf {
		fmt.Println("data", data, string(b))
	}
}
