package main

import "os"

func main() {
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString("hello")

	bytes := []byte("gogo")
	f.Write(bytes)
}
