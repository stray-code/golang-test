package main

import "fmt"

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermissions = false

	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}

	if age2 := 20; age2 >= 18 {
		fmt.Println("person is an adult", age2)
	} else {
		fmt.Println("person is teenager", age2)
	}
}
