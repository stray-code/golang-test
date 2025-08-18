package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	IsAdmin bool     `json:"isAdmin"`
	Emails  []string `json:"emails"`
}

func main() {
	user := User{
		Name:    "Kasumi",
		Age:     21,
		IsAdmin: false,
		Emails:  []string{"kasumi@example.com", "takeshi@example.com"},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(jsonData))
}
