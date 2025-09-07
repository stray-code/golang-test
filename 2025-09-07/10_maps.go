package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)

	m["name"] = "golang"
	m["area"] = "backend"

	fmt.Println(m["name"], m["area"])
	fmt.Println(m["phone"])

	m2 := make(map[string]int)
	m2["age"] = 24
	m2["price"] = 1000

	delete(m2, "price")
	fmt.Println(m2)

	object := map[string]int{"age": 14, "weight": 34}

	_, ok := object["age"]
	if ok {
		fmt.Println("all ok", object)
	} else {
		fmt.Println("not ok", object)
	}

	map1 := map[string]string{"name": "takeshi"}
	map2 := map[string]string{"name": "takeshi"}

	fmt.Println(maps.Equal(map1, map2))
}
