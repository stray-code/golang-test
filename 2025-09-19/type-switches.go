package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%s is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("i don't know about type %T\n", v)
	}
}

func main() {
	do(3)
	do("hello")
	do(true)
}
