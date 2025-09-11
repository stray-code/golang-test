package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sourceFile, err := os.Open("source.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	distFile, err := os.Create("dist.txt")
	if err != nil {
		panic(err)
	}

	defer distFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(distFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		err2 := writer.WriteByte(b)
		if err2 != nil {
			panic(err2)
		}
	}

	writer.Flush()

	fmt.Println("written to new file successfully")
}
