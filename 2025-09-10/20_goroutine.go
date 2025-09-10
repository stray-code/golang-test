package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	wg.Done()

	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for index := range 5 {
		wg.Add(1)

		go task(index, &wg)

		// go func() {
		// 	fmt.Println(index)
		// }()
	}

	wg.Wait()

	// time.Sleep(time.Second * 2)
}
