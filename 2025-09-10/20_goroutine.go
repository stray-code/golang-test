package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {
	for index := range 5 {
		go task(index)

		go func() {
			fmt.Println(index)
		}()
	}

	time.Sleep(time.Second * 2)
}
