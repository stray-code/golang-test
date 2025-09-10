package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println(num)

		time.Sleep(time.Second)
	}
}

func main() {
	numChan := make(chan int)

	go processNum(numChan)

	for {
		numChan <- rand.Intn(100)
	}
}
